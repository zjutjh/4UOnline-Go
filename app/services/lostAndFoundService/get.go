package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetLostAndFoundById 获取指定ID的失物招领
func GetLostAndFoundById(id uint) (record models.LostAndFoundRecord, err error) {
	result := database.DB.Where("id = ?", id).First(&record)
	err = result.Error
	return record, err
}
