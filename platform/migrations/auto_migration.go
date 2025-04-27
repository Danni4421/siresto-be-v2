package migrations

import (
	"github.com/Danni4421/siresto-be-v2/app/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Auto migrate models
	err := db.AutoMigrate(
		&models.User{},
		&models.Authentication{},
	)

	if err != nil {
		return err
	}

	return nil

}
