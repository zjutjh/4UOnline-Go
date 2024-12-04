package qrcodeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// SaveQrcode 向数据库中保存一条权益码记录
func SaveQrcode(qrcode models.Qrcode) error {
	result := database.DB.Save(&qrcode)
	return result.Error
}
