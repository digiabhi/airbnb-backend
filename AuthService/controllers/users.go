package controllers

import (
	"AuthService/dto"
	"AuthService/services"
	"AuthService/utils"
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
	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJSONBody(r, &payload); jsonErr != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", jsonErr)
		return
	}
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}
	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
