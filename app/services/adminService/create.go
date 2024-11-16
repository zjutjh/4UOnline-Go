package adminService

import (
	"fmt"

	"4u-go/app/models"
	"4u-go/config/database"
	"golang.org/x/crypto/bcrypt"
)

// CreateAdminUser 创建管理员用户
func CreateAdminUser(username string, password string, userType uint, college string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user := &models.User{
		Type:      userType,
		StudentID: username,
		Password:  string(hashedPassword),
		College:   college,
	}
	res := database.DB.Create(&user)

	return user, res.Error
}
