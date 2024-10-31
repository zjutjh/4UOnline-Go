package models

import "time"

// Config 系统配置项的结构体
type Config struct {
	ID         uint      `gorm:"primaryKey"` // ID 是配置项的唯一标识
	Key        string    // Key 是配置项的键，必须唯一且不能为空
	Value      string    // Value 是配置项的值，不能为空
	UpdateTime time.Time `gorm:"comment:'设置时间';type:timestamp"` // UpdateTime 是配置项的最后更新时间
}
