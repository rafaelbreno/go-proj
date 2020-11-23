package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-proj/app/helpers"
	"go-proj/app/middlewares"
	"go-proj/app/models"
	"net/http"
)

type UserController struct {
	Auth middlewares.Auth
}

func (_ UserController) SignUp(c *gin.Context) {
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
			"user":          user,
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		},
	})
}

func (_ UserController) SignIn(c *gin.Context) {
	var input, user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := helpers.VerifyPassword(input.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var token models.Token

	err := token.SetJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		},
	})
}

func (u *UserController) Logout(c *gin.Context) {
	u.Auth.GetAuth(c)

	deleted, err := models.Redis.Del(c, u.Auth.AccessUuid).Result()

	if err != nil || deleted == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}
