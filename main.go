package main

import (
	"devstream.in/pixelated-pipeline/api"
	"devstream.in/pixelated-pipeline/config"
	"github.com/charmbracelet/log"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("could not load config", "err", err)
	}
	log.Info(conf.DatabaseConf.Source)

	router := api.NewRouter()
	router.RegisterRoutes()
	router.Start()
}
