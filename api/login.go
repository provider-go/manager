package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/global"
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
		return
	}
	if len(item.Password) < 20 {
		output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
		return
	}
	if item.Password == passwordHash {
		output.ReturnSuccessResponse(ctx, nil)
	} else {
		output.ReturnErrorResponse(ctx, 9999, "用户或密码不正确~")
	}
}

func LoginByPhone(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	phone := output.ParamToString(json["phone"])
	code := output.ParamToString(json["code"])

	// 对比数据库记录,手机号是否存在
	item, err := models.ViewManagerUserByPhone(phone)
	if err != nil {
		logger.Error("LoginByUsername", "step", "ViewManagerUserByUsername", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	if len(item.Username) < 2 {
		output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
		return
	}
	// 对比缓存种是否存在手机号验证码
	value := global.Cache.Get(phone)
	if code != value {
		output.ReturnErrorResponse(ctx, 9999, "短信验证码错误~")
		return
	}
	// 删除缓存记录
	global.Cache.Del(phone)
	output.ReturnSuccessResponse(ctx, nil)
}
