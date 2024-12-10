package models

import (
	"database/sql"
	"time"
)

// Qrcode 权益码的结构体
type Qrcode struct {
	ID           uint         `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"-"`
	DeletedAt    sql.NullTime `json:"-" gorm:"index"`
	FeedbackType uint         `json:"feedback_type"` // 反馈类型
	College      uint         `json:"college"`       // 责任部门
	Department   string       `json:"department"`    // 负责单位
	Location     string       `json:"location"`      // 投放位置
	Status       bool         `json:"status"`        // 状态(是否启用)
	Description  string       `json:"description"`   // 备注

	ScanCount     uint `json:"scan_count"`     // 扫描次数
	FeedbackCount uint `json:"feedback_count"` // 反馈次数
}
