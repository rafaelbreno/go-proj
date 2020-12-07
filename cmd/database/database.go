package database

import (
	"github.com/joho/godotenv"
)

type Connection struct {
	Redis    Redis
	Postgres Postgres
}

func (c *Connection) Innit() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	c.Postgres.SetPostgres()
	c.Redis.SetRedis()
}
