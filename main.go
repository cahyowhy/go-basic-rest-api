package main

import "go-basic-rest-api/config"

func main() {
	configApp := config.GetConfig()
	app := App{}
	app.Initialize(configApp)
	app.seedsDb()
	app.Run(":3000")
}
