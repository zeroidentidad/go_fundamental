package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"postgres-jwt-react/config"
	"postgres-jwt-react/db"
	"postgres-jwt-react/errors"
	"postgres-jwt-react/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")

//Claims struct de jwt claims
type Claims struct {
	db.User
	jwt.StandardClaims
}

// Pong prueba que la API está funcionando
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

//Iniciar email msg de restablecimiento contraseña con URL de restablecimiento
func InitiatePasswordReset(c *gin.Context) {
	var createReset db.CreateReset
	c.Bind(&createReset)
	if id, ok := checkAndRetrieveUserIDViaEmail(createReset); ok {
		link := fmt.Sprintf("%s/reset/%d", config.CLIENT_URL, id)
		// Enlace de restablecimiento se devuelve en la respuesta json con fines de prueba, ya que no se integra ningún servicio de correo electrónico
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Se envió correctamente el correo de restablecimiento a " + createReset.Email, "link": link})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "errors": "Usuario no encontrado para el email: " + createReset.Email})
	}
}

func ResetPassword(c *gin.Context) {
	var resetPassword db.ResetPassword
	c.Bind(&resetPassword)
	if ok, errStr := utils.ValidatePasswordReset(resetPassword); ok {
		password := db.CreateHashedPassword(resetPassword.Password)
		_, err := db.DB.Query(db.UpdateUserPasswordQuery, resetPassword.ID, password)
		errors.HandleErr(c, err)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "La contraseña se restableció correctamente"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "errors": errStr})
	}

}

//Create nuevo usuario
func Create(c *gin.Context) {
	var user db.Register
	c.Bind(&user)
	exists := checkUserExists(user)

	valErr := utils.ValidateUser(user, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "el email ya existe")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	db.HashPassword(&user)
	_, err := db.DB.Query(db.CreateUserQuery, user.Name, user.Password, user.Email)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Usuario creado con éxito"})
}

// Session devuelve JSON de la información del usuario
func Session(c *gin.Context) {
	user, isAuthenticated := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

// Login controller
func Login(c *gin.Context) {
	var user db.Login
	c.Bind(&user)

	row := db.DB.QueryRow(db.LoginQuery, user.Email)

	var id int
	var name, email, password, createdAt, updatedAt string

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "credenciales incorrectas"})
		return
	}

	match := db.CheckPasswordHash(user.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "credenciales incorrectas"})
		return
	}

	//tiempo de vencimiento del token -> 30 minutos
	expirationTime := time.Now().Add(30 * time.Minute)

	// Crear reclamaciones JWT, que incluyen estructura del usuario y el tiempo de vencimiento
	claims := &Claims{
		User: db.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			//tiempo de caducidad, expresado en milisegundos de Unix
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Crear cadena del token JWT
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "iniciado de sesión correcto", "user": claims.User, "token": tokenString})
}

func checkUserExists(user db.Register) bool {
	rows, err := db.DB.Query(db.CheckUserExists, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

//Devuelve -1 como ID si el usuario no existe en la tabla
func checkAndRetrieveUserIDViaEmail(createReset db.CreateReset) (int, bool) {
	rows, err := db.DB.Query(db.CheckUserExists, createReset.Email)
	if err != nil {
		return -1, false
	}
	if !rows.Next() {
		return -1, false
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, false
	}
	return id, true
}
