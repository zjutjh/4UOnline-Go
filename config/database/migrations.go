package database

import (
	"4u-go/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Config{},
		&models.User{},
		&models.Announcement{},
		&models.Activity{},
		&models.LostAndFoundRecord{},
		&models.Website{},
		&models.Collection{},
	)
}
