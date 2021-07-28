package repositories

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	"z-document/db"
	"z-document/models"
)

func TestAccountRegisterAndLogin(t *testing.T) {
	db.InitMySqlDB()
	r := NewAccountRepository()

	var user models.User
	var err error

	user, err = r.CreateUser(models.User{UserName: "__test", Password: "__test"})
	if assert.Nil(t, err) {
		assert.Equal(t, user.UserName, "__test")
		assert.Equal(t, user.Password, "__test")
	} else {
		t.Log(spew.Sdump(err))
	}

	var userLogin models.User
	userLogin, err = r.Login(models.User{UserName: "__test", Password: "__test"})
	if assert.Nil(t, err) {
		assert.Equal(t, user.ID, userLogin.ID)
		assert.Equal(t, user.UserName, userLogin.UserName)
		assert.Equal(t, user.Email, userLogin.Email)
		assert.Equal(t, user.Password, userLogin.Password)
	} else {
		t.Log(spew.Sdump(err))
	}

	err = r.DeleteUser(user)
	assert.Nil(t, err)

	userLogin, err = r.Login(models.User{UserName: "__test", Password: "__test"})
	assert.NotNil(t, err)
}
