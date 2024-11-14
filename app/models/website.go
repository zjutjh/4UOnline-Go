package models

import "time"

// Website 常用网站的结构体
type Website struct {
	ID          uint      // 网站编号
	Type        uint      // 网站类型  1-学校 2-学院 3-其他
	College     uint      // 学院ID (仅在网站类型为学院时有效)
	Title       string    // 网站名称
	Description string    // 网站简介
	URL         string    // 网站地址
	Condition   string    // 访问条件
	AuthorID    uint      // 网站发布者ID
	CreateAt    time.Time // 记录创建时间
}
