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
		// 登录接口
		Router.POST("loginByUsername", api.LoginByUsername)
		Router.POST("loginByPhone", api.LoginByPhone)
		// 日志接口
		// 菜单接口
		// 角色接口
		// 用户接口
	}
}
