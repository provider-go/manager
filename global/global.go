package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/provider-go/manager/middleware"
	"github.com/provider-go/pkg/cache"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Cache  cache.Cache
	Casbin *casbin.Enforcer
	JWT    middleware.InstanceJWT
)
