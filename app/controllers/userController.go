package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-proj/app/helpers"
	"go-proj/app/models"
	"net/http"
)

type UserController struct{}

func (_ UserController) Signup(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.PasswordConfirmation {
		err := errors.New("Both passwords must be equal")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: hash,
	}

	models.DB.Create(&user)

	var token models.Token

	err = token.SetJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}
