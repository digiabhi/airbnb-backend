package app

import (
	config "AuthService/config/env"
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
	}
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr:         app.Config.Port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on", app.Config.Port)

	return server.ListenAndServe()
}
