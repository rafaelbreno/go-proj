package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

type Redis struct {
	dbName    string
	dbPort    string
	dbConnStr string
	Conn      *redis.Client
}

func (r *Redis) SetRedis() {
	r.setCredentials()
	r.setConn()
}

func (r *Redis) setCredentials() {
	r.dbName = os.Getenv("REDIS_NAME")
	r.dbPort = os.Getenv("REDIS_PORT")
	r.dbConnStr = fmt.Sprintf("%s:%s", r.dbName, r.dbPort)
}

func (r *Redis) setConn() {
	r.Conn = redis.NewClient(&redis.Options{
		Addr:     r.dbConnStr,
		Password: "",
		DB:       0,
	})
}
