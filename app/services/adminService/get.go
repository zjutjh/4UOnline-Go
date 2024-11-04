package adminService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetUserByUsername 通过用户名获取用户
func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	res := database.DB.Where("student_id = ?", username).First(user)
	return user, res.Error
}
