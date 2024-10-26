package activityService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetActivityList 获取校园活动列表
func GetActivityList(campus uint8) (activities []models.Activity, err error) {
	db := database.DB.Order("publish_time desc")
	if campus != 4 {
		db = db.Where("campus = ?", campus)
	}
	result := db.Find(&activities)
	err = result.Error
	return activities, err
}

// GetActivityById 获取指定ID的校园活动
func GetActivityById(id uint) (activity models.Activity, err error) {
	database.DB.Where("id = ?", id).First(&activity)
	err = database.DB.Error
	return activity, err
}
