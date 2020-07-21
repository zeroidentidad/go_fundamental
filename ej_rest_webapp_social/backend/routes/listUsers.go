package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zeroidentidad/twittor/db"
)

/*ListUsers obtener lista de los usuarios DB*/
func ListUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
