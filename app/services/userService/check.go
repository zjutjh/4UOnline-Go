package userService

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser 验证用户凭证
func AuthenticateUser(user *models.User, password string) error {
	if user.Type != models.Postgraduate && user.Type != models.Undergraduate {
		return CheckLocalLogin(user, password)
	}
	return CheckLogin(user.StudentID, password)
}

// CheckLogin 用户中心登录
func CheckLogin(username, password string) error {
	return userCenterService.Login(username, password)
}

// CheckLocalLogin 本地登录
func CheckLocalLogin(user *models.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return apiException.NoThatPasswordOrWrong
	}
	return err
}
