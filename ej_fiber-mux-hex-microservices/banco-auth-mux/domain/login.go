package domain

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username  string         `db:"username"`
	ClienteID sql.NullString `db:"cliente_id"`
	Cuentas   sql.NullString `db:"numeros_cuenta"`
	Rol       string         `db:"rol"`
}

func (l Login) GenerateToken() (*string, error) {
	var claims jwt.MapClaims
	if l.Cuentas.Valid && l.ClienteID.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SECRET))
	if err != nil {
		return nil, errors.New("cannot generate token")
	}
	return &signedTokenAsString, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	cuentas := strings.Split(l.Cuentas.String, ",")
	return jwt.MapClaims{
		"cliente_id": l.ClienteID.String,
		"rol":        l.Rol,
		"username":   l.Username,
		"cuentas":    cuentas,
		"exp":        time.Now().Add(TOKEN_DURATION).Unix(),
	}
}

func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"rol":      l.Rol,
		"username": l.Username,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
