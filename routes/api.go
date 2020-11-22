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
	taskRoutes()
	authRoutes()
	// Same port as go/Dockerfile
	r.Run(":8080")
}

func authRoutes() {
	controller := controllers.UserController{}

	r.POST("/signup", controller.Signup)
}

func taskRoutes() {
	controller := controllers.TaskController{}

	r.GET("/tasks/:id", controller.Index)

	taskGroup := r.Group("/task")
	{
		taskGroup.GET("/:id", controller.Show)
		taskGroup.POST("/", controller.Store)
		taskGroup.PUT("/:id", controller.Update)
		taskGroup.DELETE("/:id", controller.Delete)
	}
}

func listRoutes() {
	controller := controllers.ListController{}
	listGroup := r.Group("/list")
	{
		listGroup.GET("/", controller.Index)
		listGroup.GET("/:id", controller.Show)
		listGroup.POST("/", controller.Store)
		listGroup.PUT("/:id", controller.Update)
		listGroup.DELETE("/:id", controller.Delete)
	}
}
