package routes

import (
	"encoding/json"
	"net/http"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*GetRelation verifica si hay relacion entre 2 usuarios */
func GetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserIDRelation = ID

	var res models.RelationResponse

	status, err := db.SearchRelation(relation)
	if err != nil || status == false {
		res.Status = false
	} else {
		res.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(res)
}
