package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Token struct {
	Token      string
	UserId     uint
	Authorized bool
	ExpireAt   int64
}

func (t *Token) SetJWT(userID uint) error {
	var minutes int
	if err := godotenv.Load(); err != nil {
		return err
	}

	minutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))

	if err != nil {
		return err
	}

	secret := os.Getenv("JWT_SECRET")

	t.UserId = userID
	t.Authorized = true
	t.ExpireAt = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = t.Authorized
	atClaims["user_id"] = t.UserId
	atClaims["expire_at"] = t.ExpireAt

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	t.Token, err = at.SignedString([]byte(secret))

	if err != nil {
		return err
	}

	return nil
}
