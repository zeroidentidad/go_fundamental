package db

import (
	"github.com/zeroidentidad/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentLogin verifica data login en DB*/
func IntentLogin(email string, password string) (models.User, bool) {
	user, encontrado, _ := CheckUserExist(email)
	if encontrado == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
