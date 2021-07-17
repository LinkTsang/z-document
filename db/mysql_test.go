package db

import (
	"fmt"
	"testing"
)

func TestMysqlConnect(t *testing.T) {
	ConnectMySql()
	db := GetMySqlDB()
	fmt.Printf("db: %+v\n", db)
}
