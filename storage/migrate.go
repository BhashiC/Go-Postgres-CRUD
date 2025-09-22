package storage

import (
	"postgrestest/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Author{}, &models.Publisher{}, &models.Book{})
}
