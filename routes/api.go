package routes

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/handler"
)

var r *gin.Engine

func Innit() {
	r = gin.Default()

	routes()

	r.Run(":8080")
}

func routes() {
	userRoutes()
}

func userRoutes() {
	userH := handler.GetUserHandlers()
	r.GET("/users", userH.GetAllUsers)
}
