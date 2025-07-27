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
	userID := r.URL.Query().Get("id")
	if userID == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	fmt.Println("UserController: GetUserByID called with ID:", userID)
	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserController: CreateUser called")
	var payload dto.CreateUserRequestDTO
	if jsonErr := utils.ReadJSONBody(r, &payload); jsonErr != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Something went wrong while creating user", jsonErr)
		return
	}
	user, err := uc.UserService.CreateUser(&payload)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusCreated, "User created successfully", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}
	if len(users) == 0 {
		utils.WriteJSONSuccessResponse(w, http.StatusOK, "No users found", nil)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Users fetched successfully", users)
}

func (uc *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	fmt.Println("UserController: DeleteUserByID called with ID:", userID)
	err := uc.UserService.DeleteUserByID(userID)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User deleted successfully", nil)
}

func (uc *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Email is required", nil)
		return
	}

	fmt.Println("UserController: GetUserByEmail called with email:", email)
	user, err := uc.UserService.GetUserByEmail(email)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user by email", err)
		return
	}
	if user == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
}
