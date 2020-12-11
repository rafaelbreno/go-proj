package handler

import (
	"go-proj/cmd/helper"
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

func (uh *UserHandlers) FindById(c *gin.Context) {
	id := helper.StrToUint(c.Param("id"))

	user, err := uh.service.FindById(id)

	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"error": err.MessageContext()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
