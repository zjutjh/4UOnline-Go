package websiteService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetWebsiteList 获取网站列表
func GetWebsiteList(websiteType uint, college uint) (websites []models.Website, err error) {
	db := database.DB.Where("type = ?", websiteType)
	if websiteType == 2 && college != 0 {
		db = db.Where("college = ?", college)
	}
	result := db.Find(&websites)
	err = result.Error
	return websites, err
}

// GetAllWebsites 获取所有网站
func GetAllWebsites() (websites []models.Website, err error) {
	result := database.DB.Find(&websites)
	err = result.Error
	return websites, err
}

// GetWebsiteById 获取指定ID的网站
func GetWebsiteById(id uint) (website models.Website, err error) {
	result := database.DB.Where("id = ?", id).First(&website)
	err = result.Error
	return website, err
}
