package qrcodeService

import (
	"4u-go/app/models"
	"4u-go/config/database"
)

// DeleteQrcodeById 通过 ID 删除一条公告
func DeleteQrcodeById(qrcodeId uint) error {
	result := database.DB.Where("id = ?", qrcodeId).Delete(&models.Qrcode{})
	return result.Error
}
