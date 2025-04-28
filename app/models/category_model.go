package models

import (
	"time"

	"gorm.io/gorm"
)

type MenuCategory struct {
	ID          int16          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
