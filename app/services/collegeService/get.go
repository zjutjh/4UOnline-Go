package collegeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetCollegeList 获取学院列表
func GetCollegeList() (colleges []models.College, err error) {
	result := database.DB.Find(&colleges)
	err = result.Error
	return colleges, err
}

// GetCollegeById 获取指定ID的学院
func GetCollegeById(id uint) (college models.College, err error) {
	result := database.DB.Where("id = ?", id).First(&college)
	err = result.Error
	return college, err
}
