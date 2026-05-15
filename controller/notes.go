package controller

import (
	"api/config"
	"api/modules"

	"github.com/gin-gonic/gin"
)

func Createnote(c *gin.Context) {
	var notes modules.Notes

	if err := c.ShouldBindJSON(&notes); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid Input"})
		return
	}

	if err := config.DB.Create(&notes).Error; err != nil {
		c.JSON(500, gin.H{"Error": "Database error"})
		return
	}

	c.JSON(201, notes)
}

func Getnotes(c *gin.Context) {
	var notes []modules.Notes

	if err := config.DB.Find(&notes).Error; err != nil {
		c.JSON(500, gin.H{"Error": "Database Eror"})
		return
	}

	c.JSON(200, notes)

}

func Getbyid(c *gin.Context) {
	id := c.Param("id")
	var notes modules.Notes
	if err := config.DB.First(&notes, id).Error; err != nil {
		c.JSON(404, gin.H{"Error": "Invalid ID"})
		return
	}

	c.JSON(200, notes)
}

func Updatenotes(c *gin.Context) {
	id := c.Param("id")

	var notes modules.Notes

	if err := config.DB.First(&notes, id).Error; err != nil {
		c.JSON(404, gin.H{"Error": "No Content Found"})
		return
	}

	var input modules.Notes

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid input"})
		return
	}

	if err := config.DB.Model(&notes).Updates(input).Error; err != nil {
		c.JSON(500, gin.H{"Error": "Database update failed"})
		return
	}
	c.JSON(200, notes)
}

func Deletenotes(c *gin.Context) {
	id := c.Param("id")

	var notes modules.Notes

	if err := config.DB.First(&notes, id).Error; err != nil {
		c.JSON(404, gin.H{"Error": "Note not found"})
		return
	}

	if err := config.DB.Delete(&notes).Error; err != nil {
		c.JSON(500, gin.H{"Error": "Delete Failed"})
		return
	}
	c.Status(204)
}
