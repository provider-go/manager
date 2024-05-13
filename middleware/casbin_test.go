package middleware

import (
	"github.com/provider-go/manager/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestName(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:13306)/pyrethrum?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	InstanceCasbin(db)

	ok, err := global.Casbin.Enforce("root", "login", "PUT")
	if err != nil {
		t.Log(err)
	}
	t.Log(ok)

	ok, err = global.Casbin.AddPolicy("admin", "/api/user", "GET", "allow")
	if err != nil {
		t.Log(err)
	}
	t.Log(ok)

	ok, err = global.Casbin.Enforce("admin", "/api/user", "GET")
	if err != nil {
		t.Log(err)
	}
	t.Log(ok)
}
