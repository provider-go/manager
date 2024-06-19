package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/api"
	"github.com/provider-go/manager/middleware"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		LoginRouter := Router.Group("login")
		{
			// 登录接口
			LoginRouter.POST("loginByUsername", api.LoginByUsername)
			LoginRouter.POST("loginByPhone", api.LoginByPhone)
			LoginRouter.POST("loginByPlugin", api.LoginByPlugin)
		}

		HomeRouter := Router.Group("home").Use(middleware.CasbinAuth())
		{
			// 角色接口
			HomeRouter.POST("role", api.CreateRole)
			HomeRouter.PUT("role", api.UpdateRole)
			HomeRouter.DELETE("role", api.DeleteRole)
			HomeRouter.GET("roles", api.ListRole)
			HomeRouter.GET("role", api.ViewRole)
			// 菜单接口
			HomeRouter.POST("menu", api.CreateMenu)
			HomeRouter.PUT("menu", api.UpdateMenu)
			HomeRouter.DELETE("menu", api.DeleteMenu)
			HomeRouter.GET("menus", api.ListMenu)
			HomeRouter.GET("menu", api.ViewMenu)
			HomeRouter.GET("allMenus", api.ListAllMenu)
			// 管理员接口
			HomeRouter.GET("current/user", api.CurrentUser)
			HomeRouter.PUT("current/refreshToken", api.RefreshToken)
			HomeRouter.POST("user", api.CreateUser)
			HomeRouter.PUT("user", api.UpdateUser)
			HomeRouter.PUT("user/resetPassword", api.ResetPassword)
			HomeRouter.PUT("user/modifyPassword", api.ModifyPassword)
			HomeRouter.DELETE("user", api.DeleteUser)
			HomeRouter.GET("users", api.ListUser)
			HomeRouter.GET("user", api.ViewUser)
			// 插件接口
			HomeRouter.GET("plugins", api.ListPlugin)
			HomeRouter.GET("plugin", api.ViewPlugin)
			// 平台信息接口
			HomeRouter.GET("platform", api.ViewPlatform)
			HomeRouter.PUT("platform", api.UpdatePlatform)
			// 日志接口
			HomeRouter.GET("logs", api.ListLog)
			HomeRouter.GET("log", api.ViewLog)
		}
		/*
			接口基本访问协议：get(获取)，post(新增)，put(修改)和delete(删除)
			get /users：列出所有用户
			get /users/id：根据id获取用户
			post /user：新增用户
			put /user/id：根据用户id更新用户
			delete /user/id：根据用户id删除用户
		*/
	}
}
