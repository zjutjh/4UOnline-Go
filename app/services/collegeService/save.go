package collegeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveCollege 向数据库中保存一个学院
func SaveCollege(college models.College) error {
	result := database.DB.Save(&college)
	return result.Error
}
