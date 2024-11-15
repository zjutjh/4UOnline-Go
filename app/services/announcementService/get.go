package announcementService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetAnnouncementList 获取公告通知列表
func GetAnnouncementList() (announcements []models.Announcement, err error) {
	result := database.DB.Order("publish_time desc").Find(&announcements)
	err = result.Error
	return announcements, err
}

// GetAnnouncementById 获取指定ID的公告
func GetAnnouncementById(id uint) (announcement models.Announcement, err error) {
	result := database.DB.Where("id = ?", id).First(&announcement)
	err = result.Error
	return announcement, err
}
