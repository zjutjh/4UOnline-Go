package trackService

import (
	"errors"
	"time"

	"4u-go/app/models"
	"4u-go/config/database"
	"gorm.io/gorm"
)

// AddCount 用于更新某一字段在某一天的增量
func AddCount(name string) error {
	day := time.Now().Unix() / 86400
	counter, err := getCounter(name, day)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//	记录不存在则创建
		counter.Name = name
		counter.Day = day
		counter.Count = 1
		return saveCounter(counter)
	}
	if err != nil {
		return err
	}
	counter.Count++
	return saveCounter(counter)
}

func getCounter(name string, day int64) (counter models.Counter, err error) {
	err = database.DB.Where("name=? AND day=?", name, day).First(&counter).Error
	return counter, err
}

func saveCounter(counter models.Counter) error {
	return database.DB.Save(&counter).Error
}
