package main

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/controllers"
	"go-proj/app/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", controllers.Index)

	r.Run(":8080")
}
