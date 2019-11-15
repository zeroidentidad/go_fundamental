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
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//user:=models.User{ID:1,Username:"Zero",Password:"xd"}
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse()

	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFoundResponse(err.Error())
	} else {
		response.Data = user
	}

	userjson, _ := json.Marshal(response)
	fmt.Fprintf(w, string(userjson))
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
