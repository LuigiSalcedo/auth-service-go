package services

import (
	"authentication/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(user *models.User) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": 2 * time.Hour,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		panic(err)
	}

	return token
}
