package userService

import (
	"crypto/sha256"
	"encoding/hex"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
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
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return apiException.ServerError
	}
	pass := hex.EncodeToString(h.Sum(nil))

	if user.Password != pass {
		return apiException.NoThatPasswordOrWrong
	}
	return nil
}
