package collegeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteCollegeById 通过 ID 删除一个学院
func DeleteCollegeById(collegeId uint) error {
	result := database.DB.Where("id = ?", collegeId).Delete(&models.College{})
	return result.Error
}
