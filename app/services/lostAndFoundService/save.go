package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveLostAndFound 向数据库中保存一条失物招领
func SaveLostAndFound(record models.LostAndFoundRecord) error {
	result := database.DB.Save(&record)
	return result.Error
}
