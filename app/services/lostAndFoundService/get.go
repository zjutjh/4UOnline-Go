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

// GetLostAndFoundContact 获取失物招领联系方式
func GetLostAndFoundContact(id uint, studentID string) (contact string, err error) {
	result, err := GetLostAndFoundById(id)
	if err != nil {
		return "", err
	} else {
		var record models.ContactViewRecord
		record.RecordID = id
		record.StudentID = studentID
		res := database.DB.Save(&record)
		if res.Error != nil {
			return "", res.Error
		} else {
			return result.Contact, nil
		}
	}
}

// GetLatestLostAndFound 获取最新失物招领
func GetLatestLostAndFound() (record models.LostAndFoundRecord, err error) {
	result := database.DB.Order("created_at desc").First(&record)
	err = result.Error
	return record, err
}

// GetUserLostAndFoundStatus 查看发布失物招领信息后的审核状态
func GetUserLostAndFoundStatus(publisher string, isProcessed uint8) (records []models.LostAndFoundRecord, err error) {
	if isProcessed != 1 {
		result := database.DB.Where("publisher = ? AND is_processed = ?", publisher, isProcessed).Order("created_at desc").Find(&records)
		err = result.Error
		return records, err
	}
	result := database.DB.Where("publisher = ? AND is_approved = ?", publisher, isProcessed).Order("created_at desc").Find(&records)
	err = result.Error
	return records, err
}
