package services

import (
	db "AuthService/db/repositories"
	"AuthService/utils"
	"fmt"
)

type UserService interface {
	GetUserByID() error
	CreateUser() error
	LoginUser() error
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

func (u *UserServiceImpl) LoginUser() error {
	response := utils.CheckPasswordHash("test-password", "$2a$10$dHr8jELwcEEVkyruwKVjxuOEUp18OwMdd55bEhvOaVJRyiUFdTWs.")
	fmt.Println("Login response:", response)
	return nil
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
