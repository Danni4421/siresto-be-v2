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

	return token, refreshToken, nil
}
