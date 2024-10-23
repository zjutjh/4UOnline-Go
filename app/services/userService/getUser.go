package userService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetUserByWechatOpenID 根据微信openid获取用户
func GetUserByWechatOpenID(openid string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			WechatOpenID: openid,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

// GetUserByStudentID 根据学号获取用户
func GetUserByStudentID(sid string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			StudentID: sid,
		},
	).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	DecryptUserKeyInfo(&user)
	return &user, nil
}

// GetUserByID 根据用户ID获取用户
func GetUserByID(id int) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}
