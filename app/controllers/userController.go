package controllers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-proj/app/helpers"
	"go-proj/app/middlewares"
	"go-proj/app/models"
	"net/http"
	"os"
	"strconv"
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

// Refresh Token
func (u *UserController) Refresh(c *gin.Context) {
	var t models.Token

	if err := godotenv.Load(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	mapToken := map[string]string{}

	if err := c.ShouldBind(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := models.Redis.Del(c, refreshUuid).Result()

		if delErr != nil && deleted != 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		//Create new pairs of refresh and access tokens
		createErr := t.SetJWT(uint(userId))
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		c.JSON(http.StatusOK, gin.H{
			"access_token":  t.AccessToken,
			"refresh_token": t.RefreshToken,
		})
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
		return
	}
}
