package services

import (
	"time"

	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (s *AuthService) Authenticate(user *models.User, password string) (string, string, error) {
	const (
		tokenDuration   = 24 * time.Hour
		refreshDuration = 7 * 24 * time.Hour
	)

	if !user.CheckPassword(password) {
		return "", "", exceptions.NewUnauthorized("Invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, time.Now().Add(tokenDuration))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(user.ID, time.Now().Add(refreshDuration))
	if err != nil {
		return "", "", err
	}

	storeToken(s.DB, user, refreshToken)

	return token, refreshToken, nil
}

func storeToken(db *gorm.DB, user *models.User, token string) error {
	auth := models.Authentication{
		UserID: user.ID,
		Token:  token,
	}

	result := db.Create(&auth)
	if result.Error != nil {
		panic(result.Error)
	}

	return nil
}

func (s *AuthService) RenewateToken(user *models.User) (string, error) {
	token, err := utils.GenerateToken(user.ID, time.Now().Add(24*time.Hour))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) InvalidateToken(userID uint) error {
	auth := models.Authentication{UserID: userID}
	result := s.DB.Where("user_id = ?", userID).Delete(&auth)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *AuthService) ValidateToken(token string) (any, error) {
	claims, err := utils.ValidateToken(token, []byte(utils.GetEnv("AUTH_SECRET", "")))
	if err != nil {
		return nil, exceptions.NewUnauthorized("Invalid token")
	}

	return claims, nil
}
