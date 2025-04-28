package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID          int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Description string         `json:"description" gorm:"type:text"`
	Categories  []MenuCategory `json:"categories" gorm:"many2many:menu_categories_junction;"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type MenuCategory struct {
	ID          int16          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description string         `json:"description" gorm:"type:text"`
	Menus       []Menu         `json:"menus" gorm:"many2many:menu_categories_junction;"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type MenuCategoryJunction struct {
	MenuID         int64          `gorm:"primaryKey"`
	MenuCategoryID int16          `gorm:"primaryKey"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
