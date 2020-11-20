package routes

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/controllers"
	"go-proj/app/models"
)

var r *gin.Engine

func Init() {
	r = gin.Default()

	models.ConnectDatabase()

	listRoutes()

	// Same port as go/Dockerfile
	r.Run(":8080")
}

func listRoutes() {
	r.GET("/list", controllers.Index)
	r.GET("/list/:id", controllers.Show)
	r.POST("/list", controllers.Store)
}
