package router

import (
	"AuthService/controllers"
	"AuthService/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.With(middlewares.JWTAuthMiddleware).Get("/profile", ur.UserController.GetUserById)
	r.With(middlewares.UserCreateRequestValidator).Post("/signup", ur.UserController.CreateUser)
	r.With(middlewares.UserLoginRequestValidator).Post("/login", ur.UserController.LoginUser)
}
