package models

import "time"

// Collection 收藏的结构体
type Collection struct {
	ID       int       `json:"id"`        // 收藏编号
	UserID   int       `json:"user_id"`   // 用户编号
	Type     uint8     `json:"type"`      // 收藏类型 1-公告 2-问答 3-网站
	ObjID    int       `json:"obj_id"`    // 对象编号
	CreateAt time.Time `json:"create_at"` // 收藏时间
}
