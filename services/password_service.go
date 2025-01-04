package services

import (
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(rawPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}
