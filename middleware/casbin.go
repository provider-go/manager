package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/global"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/output"
	"gorm.io/gorm"
	"strings"
)

const (
	// StatusSuperAdmin 超级管理员
	StatusSuperAdmin = "root"
)

func InstanceCasbin(db *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logger.Error("InstanceCasbin", "step", "NewAdapterByDB", "err", err)
	}

	m, _ := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act, eft
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
	
	[matchers]
	m = g(r.sub, p.sub) && keyMatch(r.act, p.act) && keyMatch(r.obj, p.obj)
	`)

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		logger.Error("InstanceCasbin", "step", "NewEnforcer", "err", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		logger.Error("InstanceCasbin", "step", "LoadPolicy", "err", err)
	}

	// 检查基础权限是否存在,请按需修改
	if ok, _ := e.Enforce("root", "admin/user", "POST"); !ok {
		initPolicy(e)
	}
	logger.Info("InstanceCasbin", "step", "Enforce", "res", "Casbin初始化成功~")

	return e
}

func initPolicy(e *casbin.Enforcer) {
	// add base policies
	// 注意，此处一定要填写allow或deny，否则会报错
	_, err := e.AddPolicies(
		[][]string{
			//  超级管理员 有所有操作权限
			{StatusSuperAdmin, "*", "*", "allow"},
		},
	)
	if err != nil {
		logger.Error("initPolicy", "step", "AddPolicies", "err", err)
	}

	err = e.SavePolicy()
	if err != nil {
		logger.Error("initPolicy", "step", "SavePolicy", "err", err)
	}
}

// CasbinAuth 中间件
func CasbinAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		if token == "" {
			output.ReturnErrorResponse(ctx, 9999, "token不存在~")
			ctx.Abort()
			return
		}
		// token保存在ctx里
		ctx.Set("token", token)
		claims := InitJwt(global.SecretKey).ParseToken(token)
		user, err := claims.GetSubject()
		if err != nil {
			output.ReturnErrorResponse(ctx, 9999, "token已失效~")
			ctx.Abort()
			return
		}
		// 获取用户申请的资源和方法
		ctx.Set("user", user)
		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		logger.Info("CasbinAuth", "user", user, "path", path, "method", method)

		// 使用casbin提供的函数进行权限验证
		if ok, _ := global.Casbin.Enforce(user, path, method); !ok {
			output.ReturnErrorResponse(ctx, 9999, "用户无权限~")
			ctx.Abort()
			return
		}
	}
}

// GetToken 从标头或查询参数获取访问令牌
func GetToken(ctx *gin.Context) string {
	var token string
	auth := ctx.GetHeader("Authorization")
	prefix := "Bearer "

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = auth
	}

	if token == "" {
		token = ctx.Query("accessToken")
	}

	return token
}
