package routes

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/notes", controller.Createnote)
	router.GET("/notes", controller.Getnotes)
	router.GET("/notes/:id", controller.Getbyid)
	router.PUT("/notes/:id", controller.Updatenotes)
	router.DELETE("/notes/:id", controller.Deletenotes)
}
