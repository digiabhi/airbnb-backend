package main

import (
	"AuthService/app"
	dbConfig "AuthService/config/db"
	config "AuthService/config/env"
)

func main() {

	config.Load()
	dbConfig.SetupDB()

	port := app.SetConfig()

	app := app.NewApplication(port)

	app.Run()
}
