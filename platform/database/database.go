package database

import (
	"fmt"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func init() {
	err := connectDB(&dbInstance)

	if err != nil {
		panic("failed to connect database")
	}
}

func connectDB(db **gorm.DB) error {
	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "5432")
	user := utils.GetEnv("DB_USERNAME", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "postgres")
	dbname := utils.GetEnv("DB_NAME", "postgres")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	*db = d

	return nil
}

func GetDatabase() *gorm.DB {
	return dbInstance
}
