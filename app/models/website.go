package models

import "time"

// Website 常用网站的结构体
type Website struct {
	ID          uint      `json:"id"`          // 网站编号
	Type        uint      `json:"type"`        // 网站类型  1-学校 2-学院 3-其他
	College     uint      `json:"college"`     // 学院ID (仅在网站类型为学院时有效)
	Title       string    `json:"title"`       // 网站名称
	Description string    `json:"description"` // 网站简介
	URL         string    `json:"url"`         // 网站地址
	Condition   string    `json:"condition"`   // 访问条件
	AuthorID    uint      `json:"author"`      // 网站发布者ID
	CreateAt    time.Time `json:"create_at"`   // 记录创建时间
}
