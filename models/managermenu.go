package models

import (
	"github.com/provider-go/manager/global"
	"github.com/provider-go/pkg/types"
)

type ManagerMenu struct {
	ID         int32      `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	ParentID   int32      `json:"parentId" gorm:"column:parent_id;not null;default:0;comment:父ID"`
	Type       string     `json:"type" gorm:"column:type;type:varchar(20);not null;default:'';comment:菜单类型（menu、button）"`
	Code       string     `json:"code" gorm:"column:code;type:varchar(20);not null;default:'';comment:菜单编码（每个级别唯一）"`
	Name       string     `json:"name" gorm:"column:name;type:varchar(20);not null;default:'';comment:菜单显示名称"`
	Path       string     `json:"path" gorm:"column:path;type:varchar(255);not null;comment:菜单的访问路径"`
	Method     string     `json:"method" gorm:"column:method;type:varchar(20);not null;default:'';comment:请求方法"`
	APIPath    string     `json:"apiPath" gorm:"column:api_path;type:varchar(255);not null;comment:请求地址"`
	Sequence   int        `json:"sequence" gorm:"column:sequence;type:tinyint(1);not null;default:0;comment:排序顺序（按desc排序）"`
	Status     string     `json:"status" gorm:"column:status;type:varchar(10);not null;default:'';comment:菜单状态(enabled, disabled)"`
	CreateTime types.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime types.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerMenu(parentId int32, menuType, code, name, path, method, apiPath, status string) error {
	return global.DB.Table("manager_menus").
		Create(&ManagerMenu{ParentID: parentId, Type: menuType, Code: code, Name: name, Path: path, Method: method,
			APIPath: apiPath, Status: status}).Error
}

func DeleteManagerMenu(id int32) error {
	return global.DB.Table("manager_menus").Where("id = ?", id).Delete(&ManagerMenu{}).Error
}

func UpdateManagerMenu(id, parentId int32, menuType, code, name, path, method, apiPath, sequence, status string) error {
	return global.DB.Table("manager_menus").Where("id = ?", id).Updates(map[string]interface{}{
		"parent_id": parentId,
		"type":      menuType,
		"code":      code,
		"name":      name,
		"path":      path,
		"method":    method,
		"api_path":  apiPath,
		"sequence":  sequence,
		"status":    status,
	}).Error
}

func ListManagerMenu(pageSize, pageNum int) ([]*ManagerMenu, int64, error) {
	var rows []*ManagerMenu
	//计算列表数量
	var count int64
	global.DB.Table("manager_menus").Count(&count)

	if err := global.DB.Table("manager_menus").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewManagerMenu(id int32) (*ManagerMenu, error) {
	row := new(ManagerMenu)
	if err := global.DB.Table("manager_menus").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

func ListManagerMenuByParentId(parentId int) ([]*ManagerMenu, error) {
	var rows []*ManagerMenu

	if err := global.DB.Table("manager_menus").Where("parent_id = ?", parentId).Order("sequence desc").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
