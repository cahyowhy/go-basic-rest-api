package main

import (
	"flag"
	"go-basic-rest-api/config"

	"github.com/joho/godotenv"
)

func main() {
	var env = flag.String("env", "DEV", "type environment")
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
	}

	app.Run(":3000")
}
