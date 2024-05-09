package global

import (
	"github.com/provider-go/pkg/cache"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache cache.Cache
)
