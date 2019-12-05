package handlers

import (
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

func NewUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if _, err := models.CreateUser(username, password, email); err != nil {
			fmt.Println("No se pudo crear usuario")
		} else {
			fmt.Println("Usuario creado correctamente")
		}
	}
	utils.RenderTemplate(w, "users/new", nil)
}
