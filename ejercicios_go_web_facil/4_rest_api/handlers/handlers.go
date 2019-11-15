package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Put modificar usuario\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete eliminar un usuario\n")
}
