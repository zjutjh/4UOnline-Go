package models

import "time"

// User 用户结构体
type User struct {
	ID           uint      `json:"id"`             // 用户编号
	Name         string    `json:"name"`           // 姓名
	StudentID    string    `json:"student_id"`     // 学号
	Type         uint      `json:"type"`           // 用户类型
	Password     string    `json:"password"`       // 密码  （只有管理员有密码）
	WechatOpenID string    `json:"wechat_open_id"` // 微信 OpenID
	College      string    `json:"college"`        // 学院
	Class        string    `json:"class"`          // 班级
	PhoneNum     string    `json:"phone_num"`      // 手机号码
	CreatedAt    time.Time `json:"created_at"`     // 记录创建时间
}

// 用户类型常量
const (
	Undergraduate uint = iota // 本科生
	Postgraduate              // 研究生
	CollegeAdmin              // 学院管理员
	ForU                      // ForU工作人员
	SuperAdmin                // 超级管理员
)
