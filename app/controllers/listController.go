package controllers

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/models"
	"net/http"
)

func Index(c *gin.Context) {
	var lists []models.List
	models.DB.Find(&lists)

	c.JSON(http.StatusOK, gin.H{"data": lists})
}

func Store(c *gin.Context) {
	var input models.ListInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	list := models.List{
		UserId: input.UserId,
		Title:  input.Title,
		Status: input.Status,
	}

	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func Show(c *gin.Context) {
	var list models.List

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
