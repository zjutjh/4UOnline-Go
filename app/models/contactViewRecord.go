package models

import "time"

// ContactViewRecord 联系方式查看记录的结构体
type ContactViewRecord struct {
	ID        uint      `json:"id"`
	RecordID  uint      `json:"record_id"`                         // 失物招领记录编号
	StudentID string    `json:"-"`                                 // 查看者学号
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;"` // 记录创建时间
}
