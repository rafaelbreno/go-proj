package database

import (
	"github.com/go-redis/redis/v8"
	//"go-proj/routes"
	"os"
)

func Redis() *redis.Client {
	//addrs := "go-proj-redis:6379"
	addrs := os.Getenv("REDIS_NAME") + ":" + os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addrs,
		Password: "",
		DB:       0,
	})

	return rdb
}
