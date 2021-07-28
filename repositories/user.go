package repositories

import (
	"gorm.io/gorm"
	"z-document/db"
	"z-document/models"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByID(id uint) (models.User, error)
	FindUserByUsername(username string) (models.User, error)
	Login(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetMySqlDB(),
	}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	tx := r.db.Omit("ID").Create(&user)
	return user, tx.Error
}

func (r *userRepository) FindUserByID(id uint) (models.User, error) {
	var user models.User
	tx := r.db.First(&user, id)
	return user, tx.Error
}

func (r *userRepository) FindUserByUsername(username string) (models.User, error) {
	var user models.User
	tx := r.db.First(&user, "username = ?", username)
	return user, tx.Error
}

func (r *userRepository) Login(user models.User) (models.User, error) {
	var result models.User
	tx := r.db.Where("user_name = ? AND password = ?", user.UserName, user.Password).First(&result)
	return result, tx.Error
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	tx := r.db.Save(&user)
	return user, tx.Error
}

func (r *userRepository) DeleteUser(user models.User) error {
	tx := r.db.Delete(&user)
	return tx.Error
}
