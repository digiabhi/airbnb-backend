package services

import (
	env "AuthService/config/env"
	db "AuthService/db/repositories"
	"AuthService/dto"
	"AuthService/models"
	"AuthService/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)                        // GetUserByID() returns a user by ID
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) // CreateUser() creates a new user
	GetAllUsers() ([]*models.User, error)                               // GetAllUsers() returns all users
	DeleteUserByID(id string) error
	GetUserByEmail(email string) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func (u *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching user by ID:", err)
		return nil, err
	}
	if user == nil {
		fmt.Println("User not found")
		return nil, fmt.Errorf("user not found")
	}
	fmt.Println("User fetched successfully:", user)
	return user, nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in UserService")
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	fmt.Println("Fetching all users in UserService")
	users, err := u.userRepository.GetAll()
	if err != nil {
		fmt.Println("Error fetching all users:", err)
		return nil, err
	}
	fmt.Println("All users fetched successfully:", users)
	return users, nil
}

func (u *UserServiceImpl) DeleteUserByID(id string) error {
	fmt.Println("Deleting user in UserService")
	// Convert id to int64 if necessary, assuming id is a string
	// If id is already an int64, you can skip this conversion.
	idInt, err := strconv.ParseInt(id, 10, 64)
	err = u.userRepository.DeleteByID(idInt)
	if err != nil {
		fmt.Println("Error deleting user by ID:", err)
		return err
	}
	fmt.Println("User deleted successfully")
	return nil
}

func (u *UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("Fetching user by email in UserService")
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return nil, err
	}
	if user == nil {
		fmt.Println("User not found")
		return nil, fmt.Errorf("user not found")
	}
	fmt.Println("User fetched successfully:", user)
	return user, nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}

	if user == nil {
		fmt.Println("User not found")
		return "", fmt.Errorf("user not found")
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		fmt.Println("Invalid password")
		return "", nil
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN_SECRET")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	fmt.Println("Generated JWT Token:", tokenString)
	return tokenString, nil
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
