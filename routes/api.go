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
	controller := controllers.ListController{}

	r.GET("/list", controller.Index)
	r.GET("/list/:id", controller.Show)
	r.POST("/list", controller.Store)
	r.PUT("/list/:id", controller.Update)
	r.DELETE("/list/:id", controller.Delete)
}
