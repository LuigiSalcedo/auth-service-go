package services

import (
	"authentication/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func GenerateToken(user *models.User) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		panic(err)
	}

	return token
}
