package mwTypes

import (
	"github.com/casbin/casbin/v2"
	"github.com/provider-go/manager/middleware"
)

type InstanceMiddleWare struct {
	Casbin *casbin.Enforcer
	JWT    middleware.InstanceJWT
}
