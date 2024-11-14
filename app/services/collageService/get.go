package collageService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetCollageList 获取学院列表
func GetCollageList() (collages []models.Collage, err error) {
	result := database.DB.Find(&collages)
	err = result.Error
	return collages, err
}

// GetCollageById 获取指定ID的学院
func GetCollageById(id uint) (collage models.Collage, err error) {
	database.DB.Where("id = ?", id).First(&collage)
	err = database.DB.Error
	return collage, err
}
