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

	// Serve your static folder
	router.Static("/static", "./static")

	// ALIASES: This solves the 404 errors from your logs
	// It maps the "hidden" names browsers look for to your actual file
	router.StaticFile("/favicon.ico", "./static/logo.png")
	router.StaticFile("/apple-touch-icon.png", "./static/logo.png")
	router.StaticFile("/apple-touch-icon-precomposed.png", "./static/logo.png")
	router.StaticFile("/apple-touch-icon.png", "./static/logo.png")
	router.StaticFile("/apple-touch-icon-120x120.png", "./static/logo.png")

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
