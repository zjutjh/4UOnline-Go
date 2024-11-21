package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// ReviewLostAndFound 审核失物招领
func ReviewLostAndFound(recordId uint, isApproved bool) error {
	if isApproved == true {
		result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).Updates(map[string]interface{}{"is_approved": 1, "is_processed": 2})
		return result.Error
	}
	result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).Updates(map[string]interface{}{"is_approved": 0, "is_processed": 1})
	return result.Error
}
