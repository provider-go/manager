package global

import (
	"github.com/provider-go/manager/middleware/types"
	"github.com/provider-go/pkg/cache"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache cache.Cache
	MW    *mwTypes.InstanceMiddleWare
)
