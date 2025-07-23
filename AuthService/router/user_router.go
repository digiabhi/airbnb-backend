package router

import (
	"AuthService/controllers"

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
	r.Get("/profile", ur.UserController.GetUserByID)
	r.Post("/signup", ur.UserController.CreateUser)
	r.Post("/login", ur.UserController.LoginUser)
}
