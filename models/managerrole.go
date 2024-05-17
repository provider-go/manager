package models

import (
	"github.com/provider-go/manager/global"
	"github.com/provider-go/pkg/types"
)

type ManagerRole struct {
	Id          int32      `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	Code        string     `json:"code" gorm:"column:code;type:varchar(20);not null;default:'';comment:角色编码"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(200);not null;default:'';comment:角色名称"`
	Description string     `json:"description" gorm:"column:description;type:varchar(255);not null;comment:角色描述"`
	Sequence    int        `json:"sequence" gorm:"column:sequence;type:tinyint(1);not null;default:0;comment:排序顺序（按desc排序）"`
	Status      int        `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:菜单状态:0(正常)1(禁用)"`
	CreateTime  types.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  types.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerRole(code, name, description string) error {
	return global.DB.Table("manager_roles").Create(&ManagerRole{Code: code, Name: name, Description: description}).Error
}

func DeleteManagerRole(id int32) error {
	return global.DB.Table("manager_roles").Where("id = ?", id).Delete(&ManagerRole{}).Error
}

func UpdateManagerRole(id int32, code, name, description string, sequence, status int) error {
	return global.DB.Table("manager_roles").Where("id = ?", id).Updates(map[string]interface{}{
		"code":        code,
		"name":        name,
		"description": description,
		"sequence":    sequence,
		"status":      status,
	}).Error
}

func ListManagerRole(pageSize, pageNum int) ([]*ManagerRole, int64, error) {
	var rows []*ManagerRole
	//计算列表数量
	var count int64
	global.DB.Table("manager_roles").Count(&count)

	if err := global.DB.Table("manager_roles").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewManagerRole(id int32) (*ManagerRole, error) {
	row := new(ManagerRole)
	if err := global.DB.Table("manager_roles").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
