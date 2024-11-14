package collageService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// CreateCollage 向数据库中添加一个学院
func CreateCollage(collage models.Collage) error {
	result := database.DB.Create(&collage)
	return result.Error
}
