package main

func main() {
	// router := routes.NewRouter()
	config := GetConfig()
	app := App{}
	app.Initialize(config)
	app.Run(":3000")
}