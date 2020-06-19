package routes

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zeroidentidad/twittor/db"
	_jwt "github.com/zeroidentidad/twittor/jwt"
	"github.com/zeroidentidad/twittor/models"
)

/*Email valor usado en todos los EndPoints*/
var Email string

/*IDUser es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUser string

/*ProcessToken procesar token para extraer sus valores*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	secretKey := []byte(_jwt.SecretKey)
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err == nil {
		_, encontrado, _ := db.CheckUserExist(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, encontrado, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
