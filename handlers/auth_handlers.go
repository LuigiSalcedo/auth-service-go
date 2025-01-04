package handlers

import (
	"authentication/models"
	"authentication/services"
	"github.com/gin-gonic/gin"
)

type RegisterUser struct {
	Username string `json:"username" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

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
