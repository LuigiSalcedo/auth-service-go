package config

import (
	"authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var database *gorm.DB

func InitDatabaseConnection() {
	dns := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}
	database = db
}

func MigrateUserModel() {
	err := database.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate user model")
	}
}

func DB() *gorm.DB {
	return database
}
