package main

import (
	"api/config"
	"api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.Connect()

	router := gin.Default()

	// 1. SERVE STATIC FILES
	// This maps the URL path "/static" to your physical directory "./static".
	// This is essential for the browser to find icon.png.
	router.Static("/static", "./static")

	// 2. SERVE THE ENTRY POINT
	// This handles the root URL and serves your HTML file.
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// 3. INITIALIZE API ROUTES
	// Loads your existing routes for notes, banking, or travel logic.
	routes.Routes(router)

	// 4. CONFIGURE PORT FOR RAILWAY
	// Railway provides the port dynamically via environment variables.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for local development
	}

	log.Println("Server running on port", port)

	// Start the server
	router.Run(":" + port)
}
