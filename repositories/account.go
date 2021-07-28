package repositories

import (
	"gorm.io/gorm"
	"z-document/db"
	"z-document/models"
)

type AccountRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByID(id uint) (models.User, error)
	FindUserByUsername(username string) (models.User, error)
	Login(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{
		db: db.GetMySqlDB(),
	}
}

func (r *accountRepository) CreateUser(user models.User) (models.User, error) {
	tx := r.db.Omit("ID").Create(&user)
	return user, tx.Error
}

func (r *accountRepository) FindUserByID(id uint) (models.User, error) {
	var user models.User
	tx := r.db.First(&user, id)
	return user, tx.Error
}

func (r *accountRepository) FindUserByUsername(username string) (models.User, error) {
	var user models.User
	tx := r.db.First(&user, "username = ?", username)
	return user, tx.Error
}

func (r *accountRepository) Login(user models.User) (models.User, error) {
	var result models.User
	tx := r.db.Where("user_name = ? AND password = ?", user.UserName, user.Password).First(&result)
	return result, tx.Error
}

func (r *accountRepository) UpdateUser(user models.User) (models.User, error) {
	tx := r.db.Save(&user)
	return user, tx.Error
}

func (r *accountRepository) DeleteUser(user models.User) error {
	tx := r.db.Delete(&user)
	return tx.Error
}
