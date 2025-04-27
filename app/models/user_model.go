package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(120);not null;unique"`
	Phone     string         `json:"phone" gorm:"type:varchar(15);not null;unique"`
	Password  string         `json:"-" gorm:"type:text;not null"`
	Address   string         `json:"address" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Employee struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	NIK       string    `json:"nik" gorm:"type:varchar(20);not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
