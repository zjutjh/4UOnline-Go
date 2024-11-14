package models

// College 学院的结构体
type College struct {
	ID   uint   `json:"id"`   // 学院编号
	Name string `json:"name"` // 学院名称
}
