package qrcodeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// GetQrcodeById 获取指定ID的公告
func GetQrcodeById(id uint) (qrcode models.Qrcode, err error) {
	result := database.DB.Where("id = ?", id).First(&qrcode)
	err = result.Error
	return qrcode, err
}
