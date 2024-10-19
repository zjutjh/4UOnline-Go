package models

// Website 是常用网站的结构体
type Website struct {
	ID         int    `json:"id"`         // 网站编号
	Type       uint8  `json:"type"`       // 网站类型 1-学校 2-学院 3-其他
	College    string `json:"college"`    // 学院名称 (仅当网站类型为学院时有效)
	Name       string `json:"name"`       // 网站名称
	Url        string `json:"url"`        // 网站链接
	Department string `json:"department"` // 责任单位
	Condition  string `json:"condition"`  // 访问条件
}
