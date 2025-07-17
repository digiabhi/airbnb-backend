package services

import (
	db "AuthService/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	u.userRepository.Create()
	return nil
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
