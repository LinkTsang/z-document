package db

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestMysqlConnect(t *testing.T) {
	ConnectMySql()
	db := GetMySqlDB()
	fmt.Printf("db: %+v\n", db)
}

func TestMysqlCreate(t *testing.T) {
	ConnectMySql()
	db := GetMySqlDB()
	user := &User{
		UserName: "__test",
		Email:    "__test",
		NickName: "__test",
		Password: "__test",
	}
	{
		result := db.Create(user)
		t.Logf("%+v\n", result.Error)
	}
	{
		var users []User
		result := db.Find(&users)
		t.Logf("%+v\n", result.Error)
		t.Log(spew.Sdump(users))
	}
}
