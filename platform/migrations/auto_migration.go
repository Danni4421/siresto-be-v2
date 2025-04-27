package migrations

import (
	"github.com/Danni4421/siresto-be-v2/app/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := beforeMigrattion(db)

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Authentication{},
	)

	if err != nil {
		return err
	}

	return nil
}

func beforeMigrattion(db *gorm.DB) error {
	return db.Exec(`
		DO $$ BEGIN
			CREATE TYPE user_role AS ENUM ('customer', 'manager', 'admin');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;
	`).Error
}
