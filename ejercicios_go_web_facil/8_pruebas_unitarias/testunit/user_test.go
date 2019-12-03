package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"../models"
)

var user *models.User

const (
	id           = 6
	username     = "zerox"
	password     = "password"
	passwordHash = "$2a$10$vm0Ua5kMfSc2HqwjDN6m7eiXgB7zu0K6CTrW45SLxKI0ZCq44aQ3K"
	email        = "zero@email.com"
	createdAt    = "2019-12-02"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser(username, password, email)
	if user.Username != "zerox" {
		t.Error("No es posible crear el objeto user", nil)
	}
}

func TestSave(t *testing.T) {
	user := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto user", nil)
	}
}

func randomUsername() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}

func TestUniqueUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err == nil {
		t.Error("Es posible insertar registros con usernames duplicados!")
	}
}

func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'username'", username)
	if err.Error() != message {
		t.Error("Es posible que tenga un username duplicado en la base de datos!")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUser(id)
	if !equalsUser(user) || !equalsCreatedDate(user.GetCreatedDate()) {
		t.Error("No es posible obtener el usuario")
	}
}

func equalsUser(user *models.User) bool {
	return user.Username == username && user.Email == email
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener a los usuarios")
	}
}

func equalsCreatedDate(date time.Time) bool {
	t, _ := time.Parse("2006-01-02", createdAt)
	return t == date
}

func TestPassword(t *testing.T) {
	user := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60 {
		t.Error("No es posible cifrar el password")
	}
}

/*func TestDeleteUser(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("No es posible eliminar al usuario")
	}
}*/
