package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	environmentError := godotenv.Load()

	if environmentError != nil {
		panic("Error loading application environment")
	}
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
