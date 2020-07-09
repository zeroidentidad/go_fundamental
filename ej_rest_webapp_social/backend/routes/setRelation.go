package routes

import (
	"net/http"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*SetRelation enviar registro de relacion entre usuarios*/
func SetRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserIDRelation = ID

	status, err := db.InsertRelation(relation)
	if err != nil {
		http.Error(w, "Ocurrió error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
