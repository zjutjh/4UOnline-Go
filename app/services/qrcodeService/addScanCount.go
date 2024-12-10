package qrcodeService

import (
	"4u-go/app/common/counterName"
	trackService "4u-go/app/services/trackService/counter"
)

// AddScanCount 用于更新ScanCount
func AddScanCount(id uint) error {
	//	1. 更新总量
	qrcode, err := GetQrcodeById(id)
	if err != nil {
		return err
	}
	qrcode.ScanCount++

	//	2. 更新新增量
	err = trackService.AddCount(counterName.QrcodeScan)
	if err != nil {
		return err
	}

	return SaveQrcode(qrcode)
}
