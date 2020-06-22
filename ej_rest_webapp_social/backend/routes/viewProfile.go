package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zeroidentidad/twittor/db"
)

/*ViewProfile permite extraer los valores del Perfil */
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
