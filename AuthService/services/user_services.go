package services

import (
	env "AuthService/config/env"
	db "AuthService/db/repositories"
	"AuthService/dto"
	"AuthService/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID() error
	CreateUser() error
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func (u *UserServiceImpl) GetUserByID() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "test-password"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create("username-test", "testemail@email.com", hashedPassword)
	return nil
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
		"id":    user.ID,
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
