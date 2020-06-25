package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*SendTweet permite enviar el tweet a DB */
func SendTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.SaveTweet{
		UserID:  IDUser,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió error al insertar el registro, reintentar "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
