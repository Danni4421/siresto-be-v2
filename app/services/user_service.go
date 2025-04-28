package services

import (
	"slices"

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

func (service *UserService) FindUsers() ([]*models.User, error) {
	users := []*models.User{}

	if err := service.DB.Select("id", "name", "phone", "email", "created_at").Find(&users).Omit("Password").Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserService) FindUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	if err := service.DB.First(user, id).Error; err != nil {
		return nil, exceptions.NewNotFound("User not found")
	}
	return user, nil
}

func (service *UserService) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := service.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, exceptions.NewNotFound("User not found")
	}

	return user, nil
}

func (service *UserService) VerifyUserRole(userID uint, roles []models.UserRole) error {
	user := &models.User{}

	if err := service.DB.First(user, userID).Error; err != nil {
		return exceptions.NewNotFound("User not found")
	}

	isAuthorized := slices.Contains(roles, user.Role)

	if !isAuthorized {
		return exceptions.NewForbidden("You are not authorized to access this resource")
	}

	return nil
}

func (service *UserService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	existingUser, err := service.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}

	if user.Phone != "" {
		existingUser.Phone = user.Phone
	}

	if user.Address != "" {
		existingUser.Address = user.Address
	}

	if err := service.DB.Save(existingUser).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (service *UserService) DeleteUser(id uint) error {
	user := &models.User{}
	if err := service.DB.First(user, id).Error; err != nil {
		return exceptions.NewNotFound("User not found")
	}

	if err := service.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
