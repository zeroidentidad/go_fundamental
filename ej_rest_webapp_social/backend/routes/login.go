package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/jwt"
	"github.com/zeroidentidad/twittor/models"
)

/*Login manejar proceso para acceder*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", http.StatusBadRequest)
		return
	}

	document, existe := db.IntentLogin(user.Email, user.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña incorrecto ", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general Token de sesion "+err.Error(), http.StatusBadRequest)
		return
	}

	res := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "TokenBack",
		Value:   "xd", //jwtKey
		Expires: expirationTime,
	})
	// para el front se usaria localStorage de la respuesta del backend
}
