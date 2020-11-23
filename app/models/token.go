package models

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

var ctx = context.Background()

type Token struct {
	UserId     uint
	Authorized bool
	ExpireAt   int64
	RefreshAt  int64

	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string

	Redis *redis.Client
}

func (t *Token) SetJWT(userID uint) error {
	var err error
	if err := godotenv.Load(); err != nil {
		return err
	}

	err = t.setUuid()

	if err != nil {
		return err
	}

	secret := os.Getenv("JWT_SECRET")
	refresh_secret := os.Getenv("JWT_REFRESH_SECRET")

	t.UserId = userID
	t.Authorized = true

	atClaims, rtClaims := t.setClaim()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	t.AccessToken, err = at.SignedString([]byte(secret))
	t.RefreshToken, err = rt.SignedString([]byte(refresh_secret))

	if err != nil {
		return err
	}

	err = t.setAuth()

	if err != nil {
		return err
	}

	return nil
}

func (t *Token) setUuid() error {
	minutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))

	if err != nil {
		return err
	}

	days, err := strconv.Atoi(os.Getenv("JWT_REFRESH"))

	if err != nil {
		return err
	}

	t.ExpireAt = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	t.AccessUuid = uuid.New().String()

	t.RefreshAt = time.Now().Add(time.Hour * 24 * time.Duration(days)).Unix()
	t.RefreshUuid = uuid.New().String()

	return nil
}

func (t *Token) setAuth() error {
	expire_at := time.Unix(t.ExpireAt, 0)
	refresh_at := time.Unix(t.RefreshAt, 0)
	now := time.Now()

	if err := Redis.Set(ctx, t.AccessUuid, t.UserId, expire_at.Sub(now)).Err(); err != nil {
		return err
	}

	if err := Redis.Set(ctx, t.RefreshUuid, t.UserId, refresh_at.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}

func (t *Token) setClaim() (jwt.MapClaims, jwt.MapClaims) {
	atClaims := jwt.MapClaims{}
	rtClaims := jwt.MapClaims{}

	atClaims["authorized"] = t.Authorized
	atClaims["user_id"] = t.UserId
	atClaims["expire_at"] = t.ExpireAt
	atClaims["access_uuid"] = t.AccessUuid

	rtClaims["authorized"] = t.Authorized
	rtClaims["user_id"] = t.UserId
	rtClaims["expire_at"] = t.RefreshAt
	rtClaims["refresh_uuid"] = t.RefreshUuid

	return atClaims, rtClaims
}
