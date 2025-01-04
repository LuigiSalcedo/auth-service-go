package main

import (
	"authentication/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const serverPort = 8808

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()

	config.LoadBeauty()

	fmt.Println("Server running on port", serverPort)

	log.Fatal(router.Run(fmt.Sprintf(":%d", serverPort)))
}
