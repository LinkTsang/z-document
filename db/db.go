package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"znote-server-go/config"
)

var db *gorm.DB

func ConnectMySql() {
	dsn := config.GetMySqlConfig()
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetMySqlDB() *gorm.DB {
	return db
}
