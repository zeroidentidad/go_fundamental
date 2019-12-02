package test

import (
	"testing"

	"../models"
)

var user *models.User

const (
	id           = 1
	username     = "zeroidentidad"
	password     = "password"
	passwordHash = "$2a$10$vm0Ua5kMfSc2HqwjDN6m7eiXgB7zu0K6CTrW45SLxKI0ZCq44aQ3K"
	email        = "zero@email.com"
	createdDate  = "2019-12-02"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser(username, password, email)
	if user.Username != "zeroidentidad" {
		t.Error("No es posible crear el objeto user", nil)
	}
}

func TestSave(t *testing.T) {
	user := models.NewUser("zerox", password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser("zerox2", password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto user", nil)
	}
}
