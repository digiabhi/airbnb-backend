package main

import (
	"AuthService/app"
)

func main() {
	port := app.SetConfig(":8080")

	app := app.NewApplication(port)

	app.Run()
}
