package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/encryption/sm3"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/output"
)

func LoginByUsername(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	username := output.ParamToString(json["username"])
	password := output.ParamToString(json["password"])
	// 对password进行双hash
	ripemd := sm3.NewSMThree("ripemd160")
	passwordHash := ripemd.Hash([]byte(password))
	// 对比数据库记录
	item, err := models.ViewManagerUserByUsername(username)
	if err != nil {
		logger.Error("LoginByUsername", "step", "ViewManagerUserByUsername", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	}
	if len(item.Password) < 20 {
		output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
	}
	if item.Password == passwordHash {
		output.ReturnSuccessResponse(ctx, nil)
	}
}

func LoginByPhone(ctx *gin.Context) {

}
