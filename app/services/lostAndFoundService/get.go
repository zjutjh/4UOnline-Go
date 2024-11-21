package lostAndFoundService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetLostAndFoundById 获取指定ID的失物招领
func GetLostAndFoundById(id uint) (record models.LostAndFoundRecord, err error) {
	result := database.DB.Where("id = ?", id).First(&record)
	err = result.Error
	return record, err
}

// GetLostAndFoundList 获取失物招领列表
func GetLostAndFoundList(Type bool, campus, kind uint8) (records []models.LostAndFoundRecord, err error) {
	if kind == 0 {
		result := database.DB.Where("type = ? AND campus = ?", Type, campus).Order("created_at desc").Find(&records)
		err = result.Error
		return records, err
	} else {
		result := database.DB.Where("type = ? AND campus = ? AND kind = ?", Type, campus, kind).Order("created_at desc").Find(&records)
		err = result.Error
		return records, err
	}
}
