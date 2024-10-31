package activityService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveActivity 向数据库中保存一条活动记录
func SaveActivity(activity models.Activity) error {
	result := database.DB.Save(&activity)
	return result.Error
}
