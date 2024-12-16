package models

import "time"

// LostAndFoundRecord 失物招领记录的结构体
type LostAndFoundRecord struct {
	ID           uint      `json:"id"`
	Type         bool      `json:"type"`                              // 报失/捡拾 1-报失 0-捡拾
	Name         string    `json:"name"`                              // 物品名称
	Introduction string    `json:"introduction"`                      // 物品介绍
	Campus       uint8     `json:"campus"`                            // 校区 0-其他 1-朝晖 2-屏峰 3-莫干山
	Kind         uint8     `json:"kind"`                              // 物品种类 0-全部 1-其他 2-饭卡 3-电子 4-文体 5-衣包 6-证件
	Place        string    `json:"place"`                             // 丢失或拾得地点
	Time         string    `json:"time"`                              // 丢失或拾得时间
	Imgs         string    `json:"imgs"`                              // 物品图片，多个图片以逗号分隔
	Contact      string    `json:"contact"`                           // 联系方式
	ContactWay   uint8     `json:"contact_way"`                       // 联系方式选项 1-手机号 2-qq 3-微信 4-邮箱
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"` // 发布时间
	IsProcessed  uint8     `json:"is_processed"`                      // 是否完成 0-已取消 1-已完成 2-进行中
	Publisher    string    `json:"-"`                                 // 发布者
	IsApproved   uint8     `json:"-"`                                 // 是否审核通过 0-未通过 1-已通过 2-待审核
}
