package models

import "time"

// User 用户结构体
type User struct {
	ID           uint      `json:"id"`             // 用户编号
	Name         string    `json:"name"`           // 姓名
	StudentID    string    `json:"student_id"`     // 学号
	Type         UserType  `json:"type"`           // 用户类型
	Password     string    `json:"password"`       // 密码  （只有管理员有密码）
	WechatOpenID string    `json:"wechat_open_id"` // 微信 OpenID
	Collage      string    `json:"collage"`        // 学院
	Class        string    `json:"class"`          // 班级
	PhoneNum     string    `json:"phone_num"`      // 手机号码
	CreateTime   time.Time `json:"create_time"`    // 记录创建时间
}

// UserType 用户类型
type UserType uint

// 用户类型常量
const (
	Undergraduate UserType = 0
	Postgraduate  UserType = 1
	ForU          UserType = 2 // 普通工作人员
	CollageAdmin  UserType = 3 // 学院管理员
	SuperAdmin    UserType = 4
)
