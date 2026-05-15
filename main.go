package main

import (
	"api/config"
	"api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()

	router := gin.Default()

	// Serve frontend
	router.StaticFile("/", "./static/index.html")

	// API routes
	routes.Routes(router)

	// Railway dynamic port
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
