package collageService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveCollage 向数据库中保存一个学院
func SaveCollage(collage models.Collage) error {
	result := database.DB.Save(&collage)
	return result.Error
}
