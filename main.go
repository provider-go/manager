package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/manager/global"
	"github.com/provider-go/manager/middleware"

	"github.com/provider-go/manager/router"
	"github.com/provider-go/pkg/types"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func CreatePluginAndDB(instance types.PluginNeedInstance) *Plugin {
	global.DB = instance.Mysql
	global.Cache = instance.Cache
	global.SecretKey = "SecretKey"
	global.Casbin = middleware.InstanceCasbin(instance.Mysql)
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "manager"
}
