package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func LoadCors(router *gin.Engine) {
	var origins []string
	envOrigins := os.Getenv("CORS_ORIGINS")

	if envOrigins == "" {
		origins = []string{"*"}
	}

	origins = strings.Split(envOrigins, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false,
	}))
}
