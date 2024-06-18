package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/global"
	"github.com/provider-go/manager/middleware"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/encryption"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/pkg/util"
)

func CreateUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	username := output.ParamToString(json["username"])
	name := output.ParamToString(json["name"])
	password := output.ParamToString(json["password"])
	// 对password进行sm3 hash
	passwordHash := encryption.SM3Hash(password)
	phone := output.ParamToString(json["phone"])
	remark := output.ParamToString(json["remark"])
	err := models.CreateManagerUser(username, name, passwordHash, phone, remark)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func UpdateUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	username := output.ParamToString(json["username"])
	name := output.ParamToString(json["name"])
	password := output.ParamToString(json["password"])
	// 对password进行sm3 hash
	passwordHash := encryption.SM3Hash(password)
	phone := output.ParamToString(json["phone"])
	remark := output.ParamToString(json["remark"])
	status := output.ParamToString(json["status"])
	err := models.UpdateManagerUser(id, username, name, passwordHash, phone, remark, status)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}

}

func ResetPassword(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	// 生成10位随机密码
	password := util.GetRandString(10)
	// 对password进行sm3 hash
	passwordHash := encryption.SM3Hash(password)
	err := models.UpdatePasswordManagerUser(id, passwordHash)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, password)
	}

}

func ModifyPassword(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	newPassword := output.ParamToString(json["newPassword"])
	oldPassword := output.ParamToString(json["oldPassword"])
	// 对password进行sm3 hash
	newPasswordHash := encryption.SM3Hash(newPassword)
	oldPasswordHash := encryption.SM3Hash(oldPassword)
	// 查询旧密码是否匹配
	item, err := models.ViewManagerUserById(id)
	if err != nil {
		logger.Error("ModifyPassword", "step", "ViewManagerUserById", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	if item.Password != oldPasswordHash {
		output.ReturnErrorResponse(ctx, 9999, "旧密码不匹配~")
		return
	}

	err = models.UpdatePasswordManagerUser(id, newPasswordHash)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}

}

func DeleteUser(ctx *gin.Context) {
	id := output.ParamToInt32(ctx.Query("id"))
	err := models.DeleteManagerUser(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}

}

func ListUser(ctx *gin.Context) {
	pageSize := output.ParamToInt(ctx.Query("pageSize"))
	pageNum := output.ParamToInt(ctx.Query("pageNum"))

	list, total, err := models.ListManagerUser(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}

}

func ViewUser(ctx *gin.Context) {
	id := output.ParamToInt32(ctx.Query("id"))
	row, err := models.ViewManagerUserById(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}

}

func CurrentUser(ctx *gin.Context) {
	row, err := models.ViewManagerUserByUsername(ctx.GetString("user"))
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}

}

func RefreshToken(ctx *gin.Context) {
	token := ctx.GetString("token")
	newToken := middleware.InitJwt(global.SecretKey).CreateTokenByOldToken(token)
	if len(newToken) > 10 {
		output.ReturnSuccessResponse(ctx, newToken)
	} else {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	}

}
