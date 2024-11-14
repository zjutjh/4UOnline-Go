package collageService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteCollageById 通过 ID 删除一个学院
func DeleteCollageById(collageId uint) error {
	result := database.DB.Where("id = ?", collageId).Delete(&models.Collage{})
	return result.Error
}
