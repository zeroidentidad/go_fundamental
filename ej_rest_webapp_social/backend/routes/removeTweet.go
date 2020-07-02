package routes

import (
	"net/http"

	"github.com/zeroidentidad/twittor/db"
)

/*RemoveTweet para enviar a borrar un Tweet en DB*/
func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parámetro ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "Ocurrió error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
