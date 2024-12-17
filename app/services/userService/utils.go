package userService

import (
	"4u-go/app/models"
	"github.com/zjutjh/WeJH-SDK/aesHelper"
)

// DecryptUserKeyInfo 解密用户信息
func DecryptUserKeyInfo(user *models.User) error {
	if user.PhoneNum != "" {
		slt, err := aesHelper.Decrypt(user.PhoneNum)
		if err != nil {
			return err
		}
		user.PhoneNum = slt[0 : len(slt)-len(user.StudentID)]
	}
	return nil
}

// EncryptUserKeyInfo 加密用户信息
func EncryptUserKeyInfo(user *models.User) error {
	if user.PhoneNum != "" {
		num, err := aesHelper.Encrypt(user.PhoneNum + user.StudentID)
		if err != nil {
			return err
		}
		user.PhoneNum = num
	}
	return nil
}
