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
