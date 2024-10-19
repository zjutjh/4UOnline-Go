package models

import "time"

// Website 网站的结构体
type Website struct {
	ID         int       `json:"id"`         // 网站编号
	Type       uint8     `json:"type"`       // 网站类型  1-学校 2-学院 3-其他
	Collage    string    `json:"collage"`    // 学院 (仅在网站类型为学院时有效)
	Name       string    `json:"name"`       // 网站名称
	URL        string    `json:"url"`        // 网站地址
	Condition  string    `json:"condition"`  // 访问条件
	Department string    `json:"department"` // 责任单位
	CreateAt   time.Time `json:"create_at"`  // 记录创建时间
}
