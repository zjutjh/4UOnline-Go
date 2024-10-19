package models

import "time"

// LostAndFoundRecord 失物招领记录的结构体
type LostAndFoundRecord struct {
	ID               int       `json:"id"`
	Type             bool      `json:"type"`                                // 1-失物 0-寻物
	ItemName         string    `json:"item_name"`                           // 物品名称
	Introduction     string    `json:"introduction"`                        // 物品介绍
	Campus           uint8     `json:"campus"`                              // 校区 1-朝晖 2-屏峰 3-莫干山
	Kind             string    `json:"kind"`                                // 物品种类 0其他1证件2箱包3首饰4现金5电子产品6钥匙
	LostOrFoundPlace string    `json:"lost_or_found_place"`                 // 丢失或拾得地点
	LostOrFoundTime  string    `json:"lost_or_found_time"`                  // 丢失或拾得时间
	Imgs             string    `json:"imgs"`                                // 物品图片，多个图片以逗号分隔
	Publisher        string    `json:"publisher"`                           // 发布者
	PickupPlace      string    `json:"pickup_place"`                        // 失物领取地点
	Contact          string    `json:"contact"`                             // 寻物联系方式
	PublishTime      time.Time `json:"publish_time" gorm:"type:timestamp;"` // 发布时间
	IsProcessed      bool      `json:"-"`                                   // 是否已处理
}
