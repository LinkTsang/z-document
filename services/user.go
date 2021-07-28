package services

import (
	"errors"
	"gorm.io/gorm"
	"z-document/models"
	"z-document/repositories"
)

type UserService interface {
	Login(user models.User) error
	Register(user models.User) error
	Logout(user models.User) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service userService) Login(user models.User) error {
	_, err := service.repository.Login(user)
	return err
}

func (service userService) Register(user models.User) error {
	u, err := service.repository.FindUserByUsername(user.UserName)

	if err == gorm.ErrRecordNotFound {
		_, err := service.repository.CreateUser(user)
		return err
	}

	if err == nil && u.UserName != "" {
		return errors.New("username already exists")
	}

	return err
}

func (service userService) Logout(user models.User) error {
	panic("implement me")
}
