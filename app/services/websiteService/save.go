package websiteService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveWebsite 向数据库中保存一个网站
func SaveWebsite(website models.Website) error {
	result := database.DB.Save(&website)
	return result.Error
}
