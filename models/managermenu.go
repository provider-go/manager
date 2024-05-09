package models

import (
	"github.com/provider-go/manager/global"
	"time"
)

type ManagerMenu struct {
	ID          int32               `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	Code        string              `json:"code" gorm:"column:code;type:varchar(20);not null;default:'';comment:菜单编码（每个级别唯一）"`
	Name        string              `json:"name" gorm:"column:name;type:varchar(20);not null;default:'';comment:菜单显示名称"`
	Description string              `json:"description" gorm:"column:description;type:varchar(255);not null;comment:菜单详细信息"`
	Sequence    int                 `json:"sequence" gorm:"column:sequence;type:tinyint(1);not null;default:0;comment:排序顺序（按desc排序）"`
	Type        string              `json:"type" gorm:"column:type;type:varchar(20);not null;default:'';comment:菜单类型（页面、按钮）"`
	Path        string              `json:"path" gorm:"column:path;type:varchar(255);not null;comment:菜单的访问路径"`
	Properties  string              `json:"properties" gorm:"column:properties;type:json;not null;comment:菜单属性（JSON）"`
	Status      string              `json:"status" gorm:"column:status;type:varchar(10);not null;default:'';comment:菜单状态(enabled, disabled)"`
	ParentID    string              `json:"parentId" gorm:"column:parent_id;not null;default:0;comment:父ID"`
	ParentPath  string              `json:"parentPath" gorm:"column:parent_path;type:varchar(255);not null;comment:父路径"`
	Children    *ManagerMenu        `json:"children" gorm:"-"` // 子菜单
	CreateTime  time.Time           `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  time.Time           `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
	Resources   ManagerMenuResource `json:"resources" gorm:"-"` // 菜单资源
}

func CreateManagerMenu(code, name, description string, sequence int, menuType, path, properties, status, parentId, parentPath string) error {
	return global.DB.Table("manager_menus").Select("code", "name", "description", "sequence", "type", "path",
		"properties", "status", "parent_id", "parent_path").
		Create(&ManagerMenu{Code: code, Name: name, Description: description, Sequence: sequence, Type: menuType, Path: path, Properties: properties,
			Status: status, ParentID: parentId, ParentPath: parentPath}).Error
}

func DeleteManagerMenu(id int32) error {
	return global.DB.Table("manager_menus").Where("id = ?", id).Delete(&ManagerMenu{}).Error
}

func UpdateManagerMenu(id int32, code, name, description string, sequence int, menuType, path, properties, status, parentId, parentPath string) error {
	return global.DB.Table("manager_menus").Where("id = ?", id).Updates(map[string]interface{}{
		"code":        code,
		"name":        name,
		"description": description,
		"sequence":    sequence,
		"type":        menuType,
		"path":        path,
		"properties":  properties,
		"status":      status,
		"parent_id":   parentId,
		"parent_path": parentPath,
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
