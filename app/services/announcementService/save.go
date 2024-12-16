package announcementService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveAnnouncement 向数据库中保存一条公告通知
func SaveAnnouncement(announcement models.Announcement) error {
	result := database.DB.Save(&announcement)
	return result.Error
}
