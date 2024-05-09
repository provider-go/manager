package models

import (
	"github.com/provider-go/manager/global"
	"time"
)

type ManagerLogger struct {
	Id         int32     `json:"id" gorm:"auto_increment;primary_key;comment:'主键'"`
	Level      string    `json:"level" gorm:"column:level;type:varchar(10);not null;default:'';comment:日志级别"`
	TraceID    string    `json:"traceId" gorm:"column:trace_id;type:varchar(30);not null;default:'';comment:访问 ID"`
	UserID     int32     `json:"userId" gorm:"column:user_id;not null;default:0;comment:用户ID"`
	Tag        string    `json:"tag" gorm:"column:tag;type:varchar(20);not null;default:'';comment:标签"`
	Message    string    `json:"message" gorm:"column:message;type:varchar(255);not null;comment:日志消息"`
	Stack      string    `json:"stack" gorm:"column:stack;type:text;not null;comment:错误堆栈"`
	Data       string    `json:"data" gorm:"column:data;type:text;not null;comment:日志数据"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
}

func CreateManagerLogger(level, traceId string, userId int32, tag, message, stack, data string) error {
	return global.DB.Table("manager_loggers").Select("level", "traceId", "userId", "tag", "message", "stack", "data").
		Create(&ManagerLogger{Level: level, TraceID: traceId, UserID: userId, Tag: tag, Message: message, Stack: stack, Data: data}).Error
}

func ListManagerLogger(pageSize, pageNum int) ([]*ManagerLogger, int64, error) {
	var rows []*ManagerLogger
	//计算列表数量
	var count int64
	global.DB.Table("manager_loggers").Count(&count)

	if err := global.DB.Table("manager_loggers").Order("create_time desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}
