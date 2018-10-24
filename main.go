package main

import (
	"flag"
	"go-basic-rest-api/config"

	"github.com/joho/godotenv"
)

func main() {
	var env = flag.String("env", "DEV", "type environment")
	var portrun string

	flag.Parse()

	if *env == "PROD" {
		portrun = ":80"
		godotenv.Load(".env.production")
	} else {
		portrun = ":3000"
		godotenv.Load(".env.development")
	}

	configApp := config.GetConfig()

	app := App{}
	app.Initialize(configApp)

	if *env != "PROD" {
		app.seedsDb()
	}

	app.Run(portrun)
}
