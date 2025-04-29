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
		&models.MenuCategory{},
		&models.Menu{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)

	if err != nil {
		return err
	}

	return nil
}

func beforeMigrattion(db *gorm.DB) error {
	return db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_type WHERE typname = 'user_role'
			) THEN
				CREATE TYPE user_role AS ENUM ('customer', 'manager', 'admin');
			END IF;

			IF NOT EXISTS (
				SELECT 1 FROM pg_type WHERE typname = 'transaction_status'
			) THEN
				CREATE TYPE transaction_status AS ENUM ('pending', 'settled', 'cancelled');
			END IF;
		END $$;
	`).Error
}
