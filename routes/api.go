package routes

import (
	"github.com/gin-gonic/gin"
	"go-proj/app/controllers"
	"go-proj/app/middlewares"
	"go-proj/app/models"
)

var r *gin.Engine

func GetTestRouter() *gin.Engine {
	r = gin.Default()

	allRoutes()

	models.ConnectTestDatabase()

	return r
}

func Init() {
	r = gin.Default()

	allRoutes()

	models.ConnectDatabase()

	// Same port as go/Dockerfile
	r.Run(":8080")
}

func allRoutes() {
	midd := middlewares.Auth{}

	auth := r.Group("/")

	auth.Use(midd.Auth())
	{
		taskRoutes(auth)
		listRoutes(auth)

		/* Delete method because
		 * it's deleting a
		 * Redis record :)
		**/
		u := controllers.UserController{}
		auth.DELETE("/logout", u.Logout)
		auth.POST("/refresh/token", u.Refresh)
	}

	authRoutes()
}

func authRoutes() {
	controller := controllers.UserController{}

	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
}

func taskRoutes(rt *gin.RouterGroup) {
	controller := controllers.TaskController{}

	rt.GET("/tasks/:id", controller.Index)

	taskGroup := rt.Group("/task")
	{
		taskGroup.GET("/:id", controller.Show)
		taskGroup.POST("/", controller.Store)
		taskGroup.PUT("/:id", controller.Update)
		taskGroup.DELETE("/:id", controller.Delete)
	}
}

func listRoutes(rl *gin.RouterGroup) {
	controller := controllers.ListController{}
	listGroup := rl.Group("/list")
	{
		listGroup.GET("/", controller.Index)
		listGroup.GET("/:id", controller.Show)
		listGroup.POST("/", controller.Store)
		listGroup.PUT("/:id", controller.Update)
		listGroup.DELETE("/:id", controller.Delete)
	}
}
