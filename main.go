package main

import (
	"api/config"
	"api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()

	router := gin.Default()

	router.Static("/", "./static")

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	routes.Routes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	router.Run(":" + port)
}
