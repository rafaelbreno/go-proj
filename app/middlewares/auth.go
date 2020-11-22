package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-proj/app/models"
	"net/http"
	"os"
	"strconv"
)

type Auth struct {
	headerToken string
	Token       *jwt.Token
	ctx         *gin.Context
	Claims      jwt.MapClaims
	AccessUuid  string
	UserId      uint
}

func (a Auth) Auth() gin.HandlerFunc {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		var err error
		a.ctx = c
		a.headerToken = c.GetHeader("Authorization")

		if err := a.getToken(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := a.checkToken(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			return
		}
	}
}

func (a Auth) getToken() error {
	if a.headerToken == "" {
		a.ctx.JSON(http.StatusForbidden, gin.H{"error": "Empty Token"})
		return fmt.Errorf("Empty token")
	}

	token, err := jwt.Parse(a.headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	a.Token = token

	return err
}

func (a Auth) checkToken() error {
	claims, ok := a.Token.Claims.(jwt.MapClaims)
	if err := a.getToken(); err != nil {
		a.ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if !ok && !a.Token.Valid {
		return fmt.Errorf("Token not valid!")
	}

	a.Claims = claims

	return nil
}

func (a Auth) checkMetadata() error {
	accessUuid, ok := a.Claims["access_uuid"].(string)
	if !ok {
		return fmt.Errorf("Couldn't claim index: access_uuid")
	}

	id, err := strconv.Atoi(fmt.Sprintf("%.f", a.Claims["user_id"]))
	userId := uint(id)

	if err != nil {
		return fmt.Errorf("Couldn't parse index: user_id")
	}
	a.AccessUuid = accessUuid
	a.UserId = userId

	return nil
}

func (a Auth) fetchToken() error {
	id, err := models.Redis.Get(a.ctx, a.AccessUuid).Result()
	if err != nil {
		return err
	}

	userId, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return err
	}

	if uint(userId) != a.UserId {
		return fmt.Errorf("Token mismatch!")
	}

	return nil
}
