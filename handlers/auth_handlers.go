package handlers

import (
	"authentication/models"
	"authentication/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
	Username string `json:"username" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

const (
	tokenTypen = "Bearer"
)

func Register(ctx *gin.Context) {
	var registerUser RegisterUser

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := new(models.User)
	user.Username = registerUser.Username
	user.Email = registerUser.Email
	user.Password = registerUser.Password

	if services.RegisterUser(user) {
		ctx.JSON(200, gin.H{"message": "User registered successfully"})
		return
	}

	ctx.JSON(400, gin.H{"error": "User registration failed"})
}

func Login(ctx *gin.Context) {
	var loginUser LoginUser

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(loginUser.Username, loginUser.Password)

	if errors.Is(err, services.ErrUserNotFound) {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, services.ErrInvalidPassword) {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(400, Token{Token: token, Type: tokenTypen})
}

func ValidateToken(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(400, gin.H{"error": "token not provided"})
		return
	}

	user, err := services.ValidateToken(token)

	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.Request.Header.Add("X-User-ID", fmt.Sprintf("%d", user.ID))
	ctx.Next()
}
