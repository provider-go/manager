package models

import (
	"github.com/provider-go/manager/global"
	"github.com/provider-go/pkg/types"
)

type ManagerUser struct {
	Id         int32      `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	DID        string     `json:"did" gorm:"column:did;type:varchar(50);not null;default:'';comment:用户did"`
	PubKey     string     `json:"pubkey" gorm:"column:pubkey;type:varchar(100);not null;default:'';comment:用户公钥"`
	Username   string     `json:"username" gorm:"column:username;type:varchar(50);not null;default:'';comment:登录用户名"`
	Name       string     `json:"name" gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名称"`
	Password   string     `json:"password" gorm:"column:password;type:varchar(64);not null;default:'';comment:登录密码（加密）"`
	Phone      string     `json:"phone" gorm:"column:phone;type:varchar(20);not null;default:'';comment:用户电话号码"`
	Remark     string     `json:"remark" gorm:"column:remark;type:varchar(255);not null;comment:用户备注"`
	Status     int        `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;comment:用户状态:0(正常)1(禁用)"`
	CreateTime types.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime types.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerUser(username, name, password, phone, remark string) error {
	return global.DB.Table("manager_users").
		Create(&ManagerUser{Username: username, Name: name, Password: password, Phone: phone, Remark: remark}).Error
}

func DeleteManagerUser(id int32) error {
	return global.DB.Table("manager_users").Where("id = ?", id).Delete(&ManagerUser{}).Error
}

func UpdateManagerUser(id int32, username, name, password, phone, remark, status string) error {
	return global.DB.Table("manager_users").Where("id = ?", id).Updates(map[string]interface{}{
		"username": username,
		"name":     name,
		"password": password,
		"phone":    phone,
		"remark":   remark,
		"status":   status,
	}).Error
}

func UpdatePasswordManagerUser(id int32, password string) error {
	return global.DB.Table("manager_users").Where("id = ?", id).Updates(map[string]interface{}{
		"password": password,
	}).Error
}

func ListManagerUser(pageSize, pageNum int) ([]*ManagerUser, int64, error) {
	var rows []*ManagerUser
	//计算列表数量
	var count int64
	global.DB.Table("manager_users").Count(&count)

	if err := global.DB.Table("manager_users").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewManagerUserById(id int32) (*ManagerUser, error) {
	row := new(ManagerUser)
	if err := global.DB.Table("manager_users").Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

func ViewManagerUserByUsername(username string) (*ManagerUser, error) {
	row := new(ManagerUser)
	if err := global.DB.Table("manager_users").Where("username = ?", username).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

func ViewManagerUserByPhone(phone string) (*ManagerUser, error) {
	row := new(ManagerUser)
	if err := global.DB.Table("manager_users").Where("phone = ?", phone).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
