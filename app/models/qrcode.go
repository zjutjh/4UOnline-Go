package models

import "gorm.io/gorm"

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
	gorm.Model
	FeedbackType uint   // 反馈类型
	College      uint   // 责任部门
	Department   string // 负责单位
	Location     string // 投放位置
	Status       bool   // 状态(是否启用)
	Description  string // 备注

	ScanCount     uint // 扫描次数
	FeedbackCount uint // 反馈次数
}
