package userService

import (
	"4u-go/app/config"
	"4u-go/app/models"
	"4u-go/app/utils"
)

// DecryptUserKeyInfo 解密用户信息
func DecryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.PhoneNum != "" {
		slt := utils.AesDecrypt(user.PhoneNum, key)
		user.PhoneNum = slt[0 : len(slt)-len(user.StudentID)]
	}
}

// EncryptUserKeyInfo 加密用户信息
func EncryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.PhoneNum != "" {
		user.PhoneNum = utils.AesEncrypt(user.PhoneNum+user.StudentID, key)
	}
}
