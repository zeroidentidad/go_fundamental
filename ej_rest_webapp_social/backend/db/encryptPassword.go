package db

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncryptPassword para encriptar de forma segura la password recibida */
func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
