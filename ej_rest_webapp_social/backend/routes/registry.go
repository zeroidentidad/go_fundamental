package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*Registry para crear en DB el registro de usuario */
func Registry(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en datos recibidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email es requerido", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Especificar contraseña de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := db.CheckUserExist(user.Email)
	if encontrado == true {
		http.Error(w, "Ya existe usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertUser(user)
	if err != nil {
		http.Error(w, "Ocurrió error al intentar realizar registro: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
