package controllers

import (
	"AuthService/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserController: RegisterUser called")
	uc.UserService.CreateUser()
	w.Write([]byte("User registration endpoint"))
}
