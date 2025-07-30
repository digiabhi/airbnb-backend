package app

import (
	dbConfig "AuthService/config/db"
	config "AuthService/config/env"
	"AuthService/controllers"
	db "AuthService/db/repositories"
	repo "AuthService/db/repositories"
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
	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return err
	}
	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Port,
		Handler:      router.SetupRouter(uRouter, rRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on", app.Config.Port)

	return server.ListenAndServe()
}
