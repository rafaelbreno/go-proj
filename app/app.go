package app

import (
	"go-proj/domain"
	"go-proj/routes"
)

func Innit() {
	// Setting DB cmd
	domain.Setting()

	// Listening
	routes.Innit()
}
