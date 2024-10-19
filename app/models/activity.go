package models

import "time"

// Activity 活动的结构体
type Activity struct {
	ID           int       `json:"id"`                                                 // 活动编号
	Title        string    `json:"title"`                                              // 活动标题
	Introduction string    `json:"introduction"`                                       // 活动简介
	Department   string    `json:"department"`                                         // 责任单位
	StartTime    time.Time `json:"start_time"`                                         // 活动时间
	Imgs         string    `json:"imgs"`                                               // 活动宣传图片，多个图片以逗号分隔
	Campus       uint8     `json:"campus"`                                             // 校区 1-朝晖 2-屏峰 3-莫干山
	Location     string    `json:"location"`                                           // 活动地点
	PublishTime  time.Time `json:"publish_time" gorm:"comment:'发布时间';type:timestamp;"` // 活动发布时间
}
