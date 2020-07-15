package routes

import (
	"net/http"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*RemoveRelation enviar borrar relacion entre usuarios*/
func RemoveRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var relation models.Relation
	relation.UserID = IDUser
	relation.UserIDRelation = ID

	status, err := db.DeleteRelation(relation)
	if err != nil {
		http.Error(w, "Ocurrió error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
