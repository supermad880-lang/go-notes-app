package main

import (
	"api/config"
	"api/routes"
	"os"
	"path/filepath" // Isko add karo

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()

	// Railway pe folder ka path sahi se nikalne ke liye:
	workingDir, _ := os.Getwd()
	staticPath := filepath.Join(workingDir, "static")

	// 1. Static folder mount karo
	router.Static("/static", staticPath)

	// 2. Direct files serve karo (Yeh 404 ko 200 mein badal dega)
	// Agar folder ka naam 'static' hai aur file 'icon.png' hai:
	iconFile := filepath.Join(staticPath, "icon.png")

	router.StaticFile("/favicon.ico", iconFile)
	router.StaticFile("/apple-touch-icon.png", iconFile)
	router.StaticFile("/apple-touch-icon-120x120.png", iconFile)
	router.StaticFile("/apple-touch-icon-120x120-precomposed.png", iconFile)
	router.StaticFile("/apple-touch-icon-precomposed.png", iconFile)

	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(staticPath, "index.html"))
	})

	routes.Routes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
