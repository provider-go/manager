package models

import (
	"github.com/provider-go/manager/global"
	"time"
)

type ManagerMenuResource struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	MenuId     int32     `json:"menuId" gorm:"column:menu_id;not null;default:0;comment:菜单ID"`
	Method     string    `json:"method" gorm:"column:method;type:varchar(20);not null;default:'';comment:HTTP方法"`
	Path       string    `json:"path" gorm:"column:path;type:varchar(255);not null;comment:接口请求地址"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerMenuResource(menuId int32, method, path string) error {
	return global.DB.Table("manager_menu_resources").Select("menu_id", "method", "path").
		Create(&ManagerMenuResource{MenuId: menuId, Method: method, Path: path}).Error
}

func DeleteManagerMenuResource(id int32) error {
	return global.DB.Table("manager_menu_resources").Where("id = ?", id).Delete(&ManagerMenuResource{}).Error
}

func ListManagerMenuResource(pageSize, pageNum int) ([]*ManagerMenuResource, int64, error) {
	var rows []*ManagerMenuResource
	//计算列表数量
	var count int64
	global.DB.Table("manager_menu_resources").Count(&count)

	if err := global.DB.Table("manager_menu_resources").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}
