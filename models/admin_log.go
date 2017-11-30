package models

import (
	"time"
)

type AdminLog struct {
	Id        int       `gorm:"column:id;primary_key"`
	Uid       int       `gorm:"column:uid;default:'2'"`
	Ip        string    `gorm:"column:ip;size:15"`
	Browser   string    `gorm:"column:browser;size:50"`
	LoginTime time.Time `gorm:"column:login_time;type(timestamp);auto_now_add;default:'CURRENT_TIMESTAMP'"`
}

func (AdminLog) TableName() string {
	return "admin_log"
}
