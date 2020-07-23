package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zeroidentidad/twittor/db"
)

/*ViewTweetsFollowers leer tweets de todos los seguidores*/
func ViewTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.ReadTweetsFollowers(IDUser, pagina)
	if correcto == false {
		http.Error(w, "Error al leer tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
