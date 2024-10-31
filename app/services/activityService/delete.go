package activityService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteActivityById 通过 ID 删除一条活动
func DeleteActivityById(activityId uint) error {
	result := database.DB.Where("id = ?", activityId).Delete(&models.Activity{})
	return result.Error
}
