package app

import (
	"go-proj/cmd/logger"
	"go-proj/domain"
	"go-proj/routes"
)

func Innit() {
	// Setting App's logger
	logger.Innit()

	// Setting DB cmd
	domain.Setting()

	// Listening
	routes.Innit()
}
