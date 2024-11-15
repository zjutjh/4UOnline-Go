package activityService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetActivityList 获取校园活动列表
func GetActivityList() (activities []models.Activity, err error) {
	result := database.DB.Order("start_time desc").Find(&activities)
	err = result.Error
	return activities, err
}

// GetActivityById 获取指定ID的校园活动
func GetActivityById(id uint) (activity models.Activity, err error) {
	result := database.DB.Where("id = ?", id).First(&activity)
	err = result.Error
	return activity, err
}
