package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// ReviewLostAndFound 审核失物招领
func ReviewLostAndFound(recordId uint, isApproved bool) error {
	if isApproved == true {
		result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).Update("is_approved", 1)
		return result.Error
	}
	result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).Update("is_approved", 0)
	return result.Error
}
