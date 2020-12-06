package handler

import (
	"go-proj/domain"
	"go-proj/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service service.UserService
}

func GetUserHandlers() UserHandlers {
	return UserHandlers{service.NewUserService(domain.NewUserRepositoryStub())}
}

func (_ *UserHandlers) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Ahaha": "aloalo"})
}
