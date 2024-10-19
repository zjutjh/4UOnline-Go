package models

import "time"

// Activity 是校园活动的结构体
type Activity struct {
	ID           int       `json:"id"`                                                  // 活动编号
	Name         string    `json:"name"`                                                // 活动标题
	Imgs         string    `json:"imgs"`                                                // 活动图片 (多个图片以逗号分隔)
	Introduction string    `json:"introduction"`                                        // 活动简介
	Campus       uint8     `json:"campus"`                                              // 活动校区 1朝晖 2屏峰 3莫干山
	ActivityTime time.Time `json:"activity_time" gorm:"comment:'活动时间';type:timestamp;"` // 活动时间
	Place        string    `json:"place"`                                               // 活动地点
	Department   string    `json:"department"`                                          // 责任单位
	PublishTime  time.Time `json:"publish_time" gorm:"comment:'发布时间';type:timestamp;"`  // 发布时间
}
