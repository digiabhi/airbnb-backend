package router

import (
	"AuthService/controllers"
	"AuthService/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	//chiRouter.Use(middlewares.RequestLogger)

	chiRouter.Use(middleware.Logger)
	//chiRouter.Use(middlewares.RateLimitMiddleware)
	chiRouter.Get("/ping", controllers.PingHandler)

	chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.in", "/fakestoreservice"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}
