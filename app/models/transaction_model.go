package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSettled   TransactionStatus = "settled"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

type Transaction struct {
	ID                 int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID             uint                `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	User               User                `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TotalPrice         int                 `json:"total_price" gorm:"not null"`
	TransactionDetails []TransactionDetail `json:"transaction_detail" gorm:"foreignKey:TransactionID;references:ID"`
	Status             TransactionStatus   `json:"status" gorm:"type:transaction_status;default:pending"`
	CreatedAt          time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}

type TransactionDetail struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	TransactionID int         `json:"transaction_id" gorm:"foreignKey:TransactionID;references:ID"`
	Transaction   Transaction `json:"-" gorm:"foreignKey:TransactionID;references:ID"`
	MenuID        int         `json:"menu_id" gorm:"foreignKey:MenuID;references:ID"`
	Menu          Menu        `json:"menu" gorm:"foreignKey:MenuID;references:ID"`
	Quantity      int         `json:"quantity" gorm:"not null"`
	TotalPrice    float64     `json:"total_price" gorm:"not null"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     time.Time   `json:"deleted_at" gorm:"index"`
}
