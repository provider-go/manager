package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/global"
	"github.com/provider-go/manager/middleware"
	"github.com/provider-go/manager/models"
	"github.com/provider-go/pkg/encryption"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/output"
)

func LoginByUsername(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	username := output.ParamToString(json["username"])
	password := output.ParamToString(json["password"])
	// 对password进行双hash
	passwordHash := encryption.SM3Hash(password)
	// 对比数据库记录
	item, err := models.ViewManagerUserByUsername(username)
	if err != nil {
		if err.Error() == "ErrRecordNotFound" {
			output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
			return
		}
		logger.Error("LoginByUsername", "step", "ViewManagerUserByUsername", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	if item.Password == passwordHash {
		// 生成token
		token := middleware.InitJwt(global.SecretKey).GenerateToken(username)
		output.ReturnSuccessResponse(ctx, token)
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
		if err.Error() == "ErrRecordNotFound" {
			output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
			return
		}
		logger.Error("LoginByUsername", "step", "ViewManagerUserByUsername", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
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
	// 生成token
	token := middleware.InitJwt(global.SecretKey).GenerateToken(item.Username)
	output.ReturnSuccessResponse(ctx, token)
}

func LoginByPlugin(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pluginToken := output.ParamToString(json["pluginToken"])
	claims := middleware.InitJwt(global.SecretKey).ParseToken(pluginToken)
	did, err := claims.GetSubject()
	if err != nil {
		logger.Error("LoginByPlugin", "step", "GetSubject", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "token解析错误~")
		return
	}
	// 对比数据库记录
	item, err := models.ViewManagerUserByPlugin(did)
	if err != nil {
		if err.Error() == "ErrRecordNotFound" {
			output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
			return
		}
		logger.Error("LoginByPlugin", "step", "ViewManagerUserByPlugin", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	// 生成token
	token := middleware.InitJwt(global.SecretKey).GenerateToken(item.Username)
	output.ReturnSuccessResponse(ctx, token)
}
