package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// ApproveLostAndFound 审核通过失物招领
func ApproveLostAndFound(recordId uint) error {
	result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).
		Updates(map[string]any{"is_approved": 1, "is_processed": 2})
	return result.Error
}

// RejectLostAndFound 审核拒绝失物招领
func RejectLostAndFound(recordId uint) error {
	result := database.DB.Model(&models.LostAndFoundRecord{}).Where("id = ?", recordId).
		Updates(map[string]any{"is_approved": 0, "is_processed": 1})
	return result.Error
}
