package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	NickName string
	Password string
}
