package services

import (
	db "AuthService/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserByID() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func (u *UserServiceImpl) GetUserByID() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetByID()
	return nil
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
