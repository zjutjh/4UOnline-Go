package models

import "time"

// Announcement 是公告的结构体
type Announcement struct {
	ID          int       `json:"id"`                                                 // 公告编号
	Title       string    `json:"title"`                                              // 公告标题
	Content     string    `json:"content"`                                            // 公告内容
	PublishTime time.Time `json:"publish_time" gorm:"comment:'发布时间';type:timestamp;"` // 公告发布时间
}
