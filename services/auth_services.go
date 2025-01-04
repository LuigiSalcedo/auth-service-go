package services

import (
	"authentication/config"
	"authentication/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) bool {
	db := config.DB()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	result := db.Create(user)

	if result.Error == nil {
		return true
	}

	return false
}
