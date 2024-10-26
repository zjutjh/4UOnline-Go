package models

import "time"

// Activity 活动的结构体
type Activity struct {
	ID           uint      // 活动编号
	Title        string    // 活动标题
	Introduction string    // 活动简介
	Department   string    // 责任单位
	StartTime    time.Time // 活动开始时间
	EndTime      time.Time // 活动结束时间
	Imgs         string    // 活动宣传图片，多个图片以逗号分隔
	Campus       uint8     // 校区 1-朝晖 2-屏峰 3-莫干山
	Location     string    // 活动地点
	PublishTime  time.Time // 活动发布时间
	AuthorID     uint      // 活动发布者ID
}
