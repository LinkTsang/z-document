package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"z-document/config"
)

var db *gorm.DB

func InitMySqlDB() {
	dsn := config.GetMySqlConfig()
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := conn.DB()
	sqlDB.SetMaxIdleConns(8)
	sqlDB.SetMaxOpenConns(64)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db = conn
}

func GetMySqlDB() *gorm.DB {
	return db
}
