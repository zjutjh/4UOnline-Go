package announcementService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteAnnouncementById 通过 ID 删除一条公告
func DeleteAnnouncementById(announcementId uint) error {
	result := database.DB.Where("id = ?", announcementId).Delete(&models.Announcement{})
	return result.Error
}
