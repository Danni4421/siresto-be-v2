package utils

import (
	"time"

	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(GetEnv("AUTH_SECRET", ""))

	if len(jwtSecret) == 0 {
		panic("You must set your secret key for authentication")
	}
}

func GenerateToken(userId uint, expiredTime time.Time) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     expiredTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string, secretKey []byte) (any, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.NewUnauthorized("Invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		return parsedToken.Claims, nil
	}

	return nil, err
}
