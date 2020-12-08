package domain

import (
	"go-proj/cmd"
	"go-proj/cmd/database"
)

var Conn database.Connection

func Setting() {
	setConn()

	migrate()
}

// Stablishing DB conn
func setConn() {
	Conn = cmd.GetConn()
}

// Migrating Domain/Models
func migrate() {
	Conn.Postgres.Conn.AutoMigrate(&User{})
}
