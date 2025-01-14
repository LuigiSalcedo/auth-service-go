package main

import (
	"authentication/config"
	"authentication/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const serverPort = 8808

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	config.LoadEnvironmentVariables()
	config.InitDatabaseConnection()
	config.MigrateUserModel()
	config.LoadBeauty()
	config.LoadCors(router)

	router.POST("auth/register", handlers.Register)
	router.POST("auth/login", handlers.Login)
	router.GET("auth/validate", handlers.ValidateToken)

	fmt.Println("Server running on port", serverPort)

	log.Fatal(router.Run(fmt.Sprintf(":%d", serverPort)))
}
