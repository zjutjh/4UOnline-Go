package models

import "time"

// LostAndFoundRecord 失物招领记录的结构体
type LostAndFoundRecord struct {
	ID           uint      `json:"id"`
	Type         bool      `json:"type"`                              // 1-失物 0-寻物
	Name         string    `json:"name"`                              // 物品名称
	Introduction string    `json:"introduction"`                      // 物品介绍
	Campus       uint8     `json:"campus"`                            // 校区 1-朝晖 2-屏峰 3-莫干山
	Kind         uint8     `json:"kind"`                              // 物品种类 0其他1证件2箱包3首饰4现金5电子产品6钥匙
	Place        string    `json:"place"`                             // 丢失或拾得地点
	Time         string    `json:"time"`                              // 丢失或拾得时间
	Imgs         string    `json:"imgs"`                              // 物品图片，多个图片以逗号分隔
	Publisher    string    `json:"publisher"`                         // 发布者
	PickupPlace  string    `json:"pickup_place"`                      // 失物领取地点
	Contact      string    `json:"contact"`                           // 联系方式
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"` // 发布时间
	IsProcessed  bool      `json:"-"`                                 // 是否已处理 0-未处理 1-已处理
}
