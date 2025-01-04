package config

import (
	"authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func InitDatabaseConnection() {
	dns := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
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
