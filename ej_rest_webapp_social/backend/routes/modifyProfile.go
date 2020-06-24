package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*ModifyProfile modifica el perfil de usuario */
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = db.ModifyRegistry(user, IDUser)
	if err != nil {
		http.Error(w, "Ocurrión error al intentar modificar registro. Reintentar "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
