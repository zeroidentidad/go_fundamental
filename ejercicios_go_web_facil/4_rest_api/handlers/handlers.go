package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{ID: 1, Username: "Zero", Password: "xd"}
	w.Header().Set("Content-Type", "application/json")
	userjson, _ := json.Marshal(user)
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
