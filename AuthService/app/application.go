package app

import (
	config "AuthService/config/env"
	"AuthService/controllers"
	db "AuthService/db/repositories"
	"AuthService/router"
	"AuthService/services"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for the application.
type Config struct {
	Port string
}

type Application struct {
	Config Config
	Store  db.Storage
}

// Constructor for setting the configuration of the application.
func SetConfig() Config {
	port := config.GetString("PORT", ":3000")
	return Config{
		Port: port,
	}
}

// NewApplication creates a new Application instance with the provided configuration.
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store:  *db.NewStorage(),
	}
}

func (app *Application) Run() error {

	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Port,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on", app.Config.Port)

	return server.ListenAndServe()
}
