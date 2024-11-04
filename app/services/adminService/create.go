package adminService

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"4u-go/app/models"
	"4u-go/config/database"
)

// CreateAdminUser 创建管理员用户
func CreateAdminUser(username string, password string, userType models.UserType, college string) (*models.User, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(password)); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	pass := hex.EncodeToString(h.Sum(nil))
	user := &models.User{
		Type:       userType,
		StudentID:  username,
		Password:   pass,
		College:    college,
		CreateTime: time.Now(),
	}
	res := database.DB.Create(&user)

	return user, res.Error
}
