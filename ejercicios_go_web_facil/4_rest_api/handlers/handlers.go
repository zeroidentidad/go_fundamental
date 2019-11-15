package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse(w)

	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFoundResponse(err.Error())
	} else {
		response.Data = user
	}

	response.Send()
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post nuevo usuario\n")
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Put modificar usuario\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete eliminar un usuario\n")
}
