package models

import (
	"database/sql"
	"time"
)

// FeedbackType
const (
	FbActivities     uint = iota // 校园活动
	FbDiningAndShops             // 食堂及商铺
	FbDormitories                // 宿舍
	FbAcademic                   // 教学服务（选课、转专业等）
	FbFacilities                 // 校园设施
	FbClassrooms                 // 教室
	FbLibrary                    // 图书馆
	FbTransportation             // 交通
	FbSecurity                   // 安保
	FbHealthCare                 // 医疗服务
	FbPolicies                   // 学院相关政策（如综测等）
	FbOthers                     // 其他服务
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
