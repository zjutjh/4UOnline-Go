package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteLostAndFoundById 通过 ID 撤回一条失物招领
func DeleteLostAndFoundById(recordId uint) error {
	result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).Update("is_processed", 0)
	return result.Error
}
