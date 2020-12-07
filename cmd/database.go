package cmd

import (
	"go-proj/cmd/database"
)

func GetConn() database.Connection {
	conn := new(database.Connection)

	conn.Innit()

	return *conn
}
