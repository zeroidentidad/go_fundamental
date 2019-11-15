package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todos los usuarios\n")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get de un usuario\n")
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
