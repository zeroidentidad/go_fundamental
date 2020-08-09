package controller

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware comprueba si token es válido, ver godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func AuthMiddleware(c *gin.Context, jwtKey []byte) (jwt.MapClaims, bool) {
	// obtener el token de sesión de las cookies de solicitudes
	ck, err := c.Request.Cookie("token")
	fmt.Println(ck, "cookie")
	if err != nil {
		fmt.Print(err)
		return nil, false
	}

	// Obtener la cadena JWT de la cookie
	tokenString := ck.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de autenticacion inesperado: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}
