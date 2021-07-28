package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID  uint
	Title   string
	Content string
}
