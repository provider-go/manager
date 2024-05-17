package models

import (
	"github.com/provider-go/manager/global"
	"github.com/provider-go/pkg/types"
)

type ManagerRoleMenu struct {
	Id         int32      `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	RoleID     int32      `json:"roleId" gorm:"column:role_id;not null;default:0;comment:角色ID"`
	MenuID     int32      `json:"menuId" gorm:"column:menu_id;not null;default:0;comment:菜单ID"`
	CreateTime types.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime types.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerRoleMenu(roleId, menuId int32) error {
	return global.DB.Table("manager_role_menus").Select("role_id", "menu_id").
		Create(&ManagerRoleMenu{RoleID: roleId, MenuID: menuId}).Error
}

func DeleteManagerRoleMenu(id int32) error {
	return global.DB.Table("manager_role_menus").Where("id = ?", id).Delete(&ManagerRoleMenu{}).Error
}

func UpdateManagerRoleMenu(id, roleId, menuId int32) error {
	return global.DB.Table("manager_role_menus").Where("id = ?", id).Updates(map[string]interface{}{
		"role_id": roleId,
		"menu_id": menuId,
	}).Error
}

func ListManagerRoleMenu(pageSize, pageNum int) ([]*ManagerRoleMenu, int64, error) {
	var rows []*ManagerRoleMenu
	//计算列表数量
	var count int64
	global.DB.Table("manager_role_menus").Count(&count)

	if err := global.DB.Table("manager_role_menus").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewManagerRoleMenu(id int32) (*ManagerRoleMenu, error) {
	row := new(ManagerRoleMenu)
	if err := global.DB.Table("manager_role_menus").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
