package services

import (
	"z-document/models"
	"z-document/repositories"
)

type AccountService interface {
	Login(user models.User) error
	Register(user models.User) error
}

type accountService struct {
	repository repositories.AccountRepository
}

func (s accountService) Login(user models.User) error {
	panic("implement me")
}

func (s accountService) Register(user models.User) error {
	panic("implement me")
}

func NewAccountService(accountRepository repositories.AccountRepository) AccountService {
	return &accountService{
		repository: accountRepository,
	}
}