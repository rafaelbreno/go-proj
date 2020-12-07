package domain

import (
	"go-proj/cmd"
	"go-proj/cmd/database"
)

var conn database.Connection

func Setting() {
	setConn()

	migrate()
}

// Stablishing DB conn
func setConn() {
	conn = cmd.GetConn()
}

// Migrating Domain/Models
func migrate() {
	conn.Postgres.Conn.AutoMigrate(&User{})
}
