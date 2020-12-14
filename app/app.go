package app

import (
	"go-proj/cmd/logger"
	"go-proj/cmd/seeder"
	"go-proj/domain"
	"go-proj/routes"
)

func Innit() {
	// Setting App's logger
	logger.Innit()

	// Setting DB cmd
	domain.Setting()

	// Seeding
	seeder.Init(domain.Conn.Postgres.Conn)

	// Listening
	routes.Innit()

}
