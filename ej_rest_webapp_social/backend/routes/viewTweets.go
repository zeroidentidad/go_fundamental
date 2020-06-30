package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zeroidentidad/twittor/db"
)

/*ViewTweets para peticion de leer los tweets */
func ViewTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(pagina)
	respuesta, correcto := db.ReadTweets(ID, page)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
