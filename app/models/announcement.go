package models

import "time"

// Announcement 公告的结构体
type Announcement struct {
	ID        uint      // 公告编号
	Title     string    // 公告标题
	Content   string    // 公告内容
	CreatedAt time.Time // 公告发布时间
	AuthorID  uint      // 公告发布者ID
}
