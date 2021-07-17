package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

var mySqlConfig string

func GetMySqlConfig() string {
	return mySqlConfig
}

func init() {
	appConfig, err := godotenv.Read("./config/config.env")
	if err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	mySqlConfig = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)
}
