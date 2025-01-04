package services

import (
	"authentication/config"
	"authentication/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
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

func Login(username string, password string) (string, error) {
	db := config.DB()
	user := new(models.User)
	db.Where("username = ?", username).First(user)

	if user.ID == 0 {
		return "", ErrUserNotFound
	}

	if !verifyPassword(password, user.Password) {
		return "", ErrInvalidPassword
	}

	return GenerateToken(user), nil
}
