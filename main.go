package main

import (
	"flag"
	"go-basic-rest-api/config"

	"github.com/joho/godotenv"
)

func main() {
	configApp := config.GetConfig()
	var env = flag.String("env", "DEV", "type environment")
	configApp.SetEnv(*env)

	if *env == "PROD" {
		godotenv.Load(".env.production")
	}

	godotenv.Load(".env.development")

	app := App{}
	app.Initialize(configApp)
	app.seedsDb()
	app.Run(":3000")
}
