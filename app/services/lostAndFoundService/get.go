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
func GetLostAndFoundList(form bool, campus, kind uint8) (records []models.LostAndFoundRecord, err error) {
	if kind == 0 {
		result := database.DB.
			Where("type =? AND campus =? AND is_processed = 2 AND is_approved = 1", form, campus).
			Order("created_at desc").
			Find(&records)
		err = result.Error
		return records, err
	}
	result := database.DB.
		Where("type = ? AND campus = ? AND kind = ? AND is_processed = 2 AND is_approved = 1", form, campus, kind).
		Order("created_at desc").
		Find(&records)
	err = result.Error
	return records, err
}

// GetLostAndFoundContact 获取失物招领联系方式
func GetLostAndFoundContact(id uint, studentID string) (contact string, err error) {
	result, err := GetLostAndFoundById(id)
	if err != nil {
		return "", err
	}
	var record models.ContactViewRecord
	record.RecordID = id
	record.StudentID = studentID
	res := database.DB.Save(&record)
	if res.Error != nil {
		return "", res.Error
	}
	return result.Contact, nil
}

// GetLatestLostAndFound 获取最新失物招领
func GetLatestLostAndFound() (record models.LostAndFoundRecord, err error) {
	result := database.DB.Where("is_processed = 2 AND is_approved = 1").Order("created_at desc").First(&record)
	err = result.Error
	return record, err
}

// GetUserLostAndFoundStatus 查看失物招领信息的状态
func GetUserLostAndFoundStatus(publisher string, status uint8) (records []models.LostAndFoundRecord, err error) {
	if status == 0 {
		result := database.DB.Where("publisher = ? AND is_processed = ?", publisher, 0).
			Order("created_at desc").
			Find(&records)
		err = result.Error
		return records, err
	} else if status == 1 {
		result := database.DB.Where("publisher = ? AND (is_approved = 0 OR is_approved = 1)", publisher).
			Order("created_at desc").
			Find(&records)
		err = result.Error
		return records, err
	}
	result := database.DB.Where("publisher = ? AND is_approved = ?", publisher, 2).
		Order("created_at desc").
		Find(&records)
	err = result.Error
	return records, err
}
