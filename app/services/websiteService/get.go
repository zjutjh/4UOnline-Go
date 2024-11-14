package websiteService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetWebsiteList 获取网站列表
func GetWebsiteList() (websites []models.Website, err error) {
	result := database.DB.Order("id desc").Find(&websites)
	err = result.Error
	return websites, err
}

// GetWebsiteById 获取指定ID的网站
func GetWebsiteById(id uint) (website models.Website, err error) {
	database.DB.Where("id = ?", id).First(&website)
	err = database.DB.Error
	return website, err
}
