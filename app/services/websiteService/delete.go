package websiteService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteWebsiteById 通过 ID 删除一个网站
func DeleteWebsiteById(websiteId uint) error {
	result := database.DB.Where("id = ?", websiteId).Delete(&models.Website{})
	return result.Error
}
