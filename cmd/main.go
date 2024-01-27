package main

import (
	"github.com/Corray333/progg/internal/app"
	"github.com/Corray333/progg/internal/config"
)

func main() {
	config.LoadConfig()
	app := app.NewApp()
	app.Run()
}
