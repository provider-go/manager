package models

import (
	"github.com/provider-go/manager/global"
	"time"
)

type ManagerUser struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	Username   string    `json:"username" gorm:"column:username;type:varchar(50);not null;default:'';comment:登录用户名"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名称"`
	Password   string    `json:"password" gorm:"column:password;type:varchar(64);not null;default:'';comment:登录密码（加密）"`
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(20);not null;default:'';comment:用户电话号码"`
	Email      string    `json:"email" gorm:"column:email;type:varchar(100);not null;comment:用户的电子邮件"`
	Remark     string    `json:"remark" gorm:"column:remark;type:varchar(255);not null;comment:用户备注"`
	Status     string    `json:"status" gorm:"column:status;type:varchar(20);not null;default:'activated';comment:用户状态（activated, freezed）"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateManagerUser(username, name, password, phone, email, remark, status string) error {
	return global.DB.Table("manager_users").Select("username", "name", "password", "phone", "email", "remark", "status").
		Create(&ManagerUser{Username: username, Name: name, Password: password, Phone: phone, Email: email, Remark: remark, Status: status}).Error
}

func DeleteManagerUser(id int32) error {
	return global.DB.Table("manager_users").Where("id = ?", id).Delete(&ManagerUser{}).Error
}

func UpdateManagerUser(id int32, username, name, password, phone, email, remark, status string) error {
	return global.DB.Table("manager_users").Where("id = ?", id).Updates(map[string]interface{}{
		"username": username,
		"name":     name,
		"password": password,
		"phone":    phone,
		"email":    email,
		"remark":   remark,
		"status":   status,
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
