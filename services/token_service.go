package services

import (
	"authentication/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
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

func ValidateToken(tokenString string) (*models.User, error) {
	if len(tokenString) < 7 {
		return nil, ErrInvalidToken
	}

	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	user := new(models.User)
	user.ID = uint(claims["id"].(float64))

	return user, nil
}
