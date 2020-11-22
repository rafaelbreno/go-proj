package models

import (
	"github.com/go-redis/redis/v8"
	"go-proj/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Redis *redis.Client

func ConnectDatabase() {
	mig := database.Conn()

	mig.AutoMigrate(&List{}, &Task{}, &User{})

	DB = mig

	Redis = database.Redis()
}
