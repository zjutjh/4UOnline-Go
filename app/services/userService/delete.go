package userService

import (
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
	"4u-go/config/database"
)

// DeleteAccount 注销账户
func DeleteAccount(user *models.User, iid string) error {
	err := userCenterService.DeleteAccount(user.StudentID, iid)
	if err != nil {
		return err
	}

	result := database.DB.Delete(user)
	return result.Error
}
