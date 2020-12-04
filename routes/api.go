package routes

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Innit() {
	r = gin.Default()

	r.Run(":8080")
}
