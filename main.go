package main

import (
	"flag"
	"go-basic-rest-api/config"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var env = flag.String("env", "DEV", "type environment")
	var port string
	flag.Parse()

	if *env == "PROD" {
		godotenv.Load(".env.production")
	} else {
		godotenv.Load(".env.development")
	}

	configApp := config.GetConfig()

	app := App{}
	app.Initialize(configApp)

	if *env != "PROD" {
		app.seedsDb()
		port = "3000"
	} else {
		//heroku use default port from env vars
		port = os.Getenv("PORT")
	}

	app.Run(":" + port)
}
