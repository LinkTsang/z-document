package services

import (
	"errors"
	"gorm.io/gorm"
	"z-document/models"
	"z-document/repositories"
)

type AccountService interface {
	Login(user models.User) error
	Register(user models.User) error
	Logout(user models.User) error
}

type accountService struct {
	repository repositories.AccountRepository
}

func NewAccountService(accountRepository repositories.AccountRepository) AccountService {
	return &accountService{
		repository: accountRepository,
	}
}

func (service accountService) Login(user models.User) error {
	_, err := service.repository.Login(user)
	return err
}

func (service accountService) Register(user models.User) error {
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

func (service accountService) Logout(user models.User) error {
	panic("implement me")
}
