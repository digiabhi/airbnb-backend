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

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserController: GetUserByID called")
	uc.UserService.GetUserByID()
	w.Write([]byte("User fetched successfully"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserController: CreateUser called")
	uc.UserService.CreateUser()
	w.Write([]byte("User Created successfully"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserController: LoginUser called")
	uc.UserService.LoginUser()
	w.Write([]byte("User logged in successfully"))
}
