package repositories

import "znote-server-go/models"

type AccountRepository interface {
	Login(user models.User) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type accountRepository struct {
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (r *accountRepository) Login(user models.User) (models.User, error) {
	panic("implement me")
}

func (r *accountRepository) CreateUser(user models.User) (models.User, error) {
	panic("implement me")
}
