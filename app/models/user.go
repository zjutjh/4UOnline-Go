package models

import "time"

// User 用户结构体
type User struct {
	ID           int       `json:"id"`             // 用户编号
	Name         string    `json:"name"`           // 姓名
	StudentID    string    `json:"student_id"`     // 学号
	Type         uint      `json:"type"`           // 用户类型  1-本科生 2-研究生
	WechatOpenID string    `json:"wechat_open_id"` // 微信 OpenID
	Collage      string    `json:"collage"`        // 学院
	Class        string    `json:"class"`          // 班级
	PhoneNum     string    `json:"phone_num"`      // 手机号码
	CreateTime   time.Time `json:"create_time"`    // 记录创建时间
}
