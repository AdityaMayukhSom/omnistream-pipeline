package main

import (
	"devstream.in/pixelated-pipeline/api"
	"devstream.in/pixelated-pipeline/config"
	"github.com/charmbracelet/log"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	err := config.LoadApplicationConfig()
	if err != nil {
		log.Fatal("could not load config", "err", err)
	}

	// db := database.Init()
	// db.Migrate()
	// db.Close()

	router := api.NewRouter()
	router.RegisterRoutes()
	router.Start()
}
