package models

import (
	"github.com/provider-go/manager/global"
	"time"
)

type ManagerUserRole struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	UserID     int32     `json:"userId" gorm:"column:user_id;not null;default:0;comment:用户ID"`
	RoleID     int32     `json:"roleId" gorm:"column:role_id;not null;default:0;comment:角色ID"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerUserRole(userId, roleId int32) error {
	return global.DB.Table("manager_role_menus").Select("user_id", "role_id").
		Create(&ManagerUserRole{UserID: userId, RoleID: roleId}).Error
}

func DeleteManagerUserRole(id int32) error {
	return global.DB.Table("manager_role_menus").Where("id = ?", id).Delete(&ManagerUserRole{}).Error
}

func UpdateManagerUserRole(id, userId, roleId int32) error {
	return global.DB.Table("manager_role_menus").Where("id = ?", id).Updates(map[string]interface{}{
		"user_id": userId,
		"role_id": roleId,
	}).Error
}

func ListManagerUserRole(pageSize, pageNum int) ([]*ManagerUserRole, int64, error) {
	var rows []*ManagerUserRole
	//计算列表数量
	var count int64
	global.DB.Table("manager_role_menus").Count(&count)

	if err := global.DB.Table("manager_role_menus").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewManagerUserRole(id int32) (*ManagerUserRole, error) {
	row := new(ManagerUserRole)
	if err := global.DB.Table("manager_role_menus").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
