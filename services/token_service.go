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

func VerifyToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	user := new(models.User)
	user.Username = claims["username"].(string)
	user.Email = claims["email"].(string)

	return user, nil
}
