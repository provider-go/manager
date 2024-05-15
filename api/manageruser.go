package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/encryption/sm3"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/pkg/util"
)

func CreateUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	username := output.ParamToString(json["username"])
	name := output.ParamToString(json["name"])
	password := output.ParamToString(json["password"])
	// 对password进行双hash
	ripemd := sm3.NewSMThree("ripemd160")
	passwordHash := ripemd.Hash([]byte(password))
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
	// 对password进行双hash
	ripemd := sm3.NewSMThree("ripemd160")
	passwordHash := ripemd.Hash([]byte(password))
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
	// 对password进行双hash
	ripemd := sm3.NewSMThree("ripemd160")
	passwordHash := ripemd.Hash([]byte(password))
	err := models.ResetPasswordManagerUser(id, passwordHash)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, password)
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
