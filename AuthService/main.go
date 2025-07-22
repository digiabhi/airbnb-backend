package main

import (
	"AuthService/app"
	
	config "AuthService/config/env"
)

func main() {

	config.Load()
	

	port := app.SetConfig()

	app := app.NewApplication(port)

	app.Run()
}
