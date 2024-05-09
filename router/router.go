package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		/*
			接口基本访问协议：get(获取)，post(新增)，put(修改)和delete(删除)
			get /users：列出所有用户
			get /users/id：根据id获取用户
			post /user：新增用户
			put /user/id：根据用户id更新用户
			delete /user/id：根据用户id删除用户
		*/
		// 角色接口
		Router.POST("role", api.CreateRole)
		Router.PUT("role", api.UpdateRole)
		Router.DELETE("role", api.DeleteRole)
		Router.GET("roles", api.ListRole)
		Router.GET("role", api.ViewRole)
		// 菜单接口
		Router.POST("menu", api.CreateMenu)
		Router.PUT("menu", api.UpdateMenu)
		Router.DELETE("menu", api.DeleteMenu)
		Router.GET("menus", api.ListMenu)
		Router.GET("menu", api.ViewMenu)
		// 管理员接口
		Router.POST("user", api.CreateUser)
		Router.PUT("user", api.UpdateUser)
		Router.DELETE("user", api.DeleteUser)
		Router.GET("users", api.ListUser)
		Router.GET("user", api.ViewUser)
		// 插件接口
		Router.GET("plugins", api.ListPlugin)
		Router.GET("plugin", api.ViewPlugin)
		// 登录接口
		Router.POST("loginByUsername", api.LoginByUsername)
		Router.POST("loginByPhone", api.LoginByPhone)
		// 平台信息接口
		Router.GET("platform", api.ViewPlatform)
		Router.PUT("platform", api.UpdatePlatform)
		// 日志接口
		Router.GET("logs", api.ListLog)
		Router.GET("log", api.ViewLog)

	}
}
