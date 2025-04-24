package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Email     string `json:"email" gorm:"type:varchar(120);not null;unique"`
	Phone     string `json:"phone" gorm:"type:char(15);not null;unique"`
	Password  string `json:"password" gorm:"type:text;not null"`
	Address   string `json:"address" gorm:"type:text"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
