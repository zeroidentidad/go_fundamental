package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zeroidentidad/twittor/models"
)

/*GenerateJWT genera datos encriptados con JWT */
func GenerateJWT(user models.User) (string, error) {

	secretKey := []byte("madafakasxd")

	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"apellidos":        user.Apellidos,
		"fecha_nacimiento": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitioweb":         user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
