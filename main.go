package main

import (
	"devstream.in/pixelated-pipeline/api"
	"devstream.in/pixelated-pipeline/config"
	"devstream.in/pixelated-pipeline/database"
	"github.com/charmbracelet/log"

	_ "devstream.in/pixelated-pipeline/docs"
)

//	@title			Pixelated Pipeline API
//	@version		1.0
//	@description	This is a sample server Petstore server hello world.
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1
func main() {
	err := config.LoadApplicationConfig()
	if err != nil {
		log.Fatal("could not load config", "err", err)
	}

	if config.ShallRunMigration() {
		db := database.Init()
		db.Migrate()
		db.Close()
	}

	router := api.NewRouter()
	router.RegisterRoutes(config.GetAllowedOrigins())
	router.Start()
}
