package services

import (
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	userExist := service.DB.Where("email = ?", user.Email).First(&models.User{})

	if userExist.RowsAffected > 0 {
		return nil, exceptions.NewNotFound("User already exists")
	}

	hashedPassword, passwordErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if passwordErr != nil {
		panic(passwordErr)
	}

	user.Password = string(hashedPassword)

	if err := service.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
