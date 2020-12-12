package main

import (
	"go-proj/app"
	"go-proj/cmd/logger"
)

func main() {

	defer logger.Log().Sync()

	app.Innit()
}
