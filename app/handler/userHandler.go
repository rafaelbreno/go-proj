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
	return UserHandlers{service.NewUserService(domain.NewUserRepositoryDB())}
}

func (uh *UserHandlers) FindAll(c *gin.Context) {
	users, _ := uh.service.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}
