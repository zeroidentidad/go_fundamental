package handlers

import (
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})

	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if _, err := models.CreateUser(username, password, email); err != nil {
			errorMessage := err.Error()
			context["Error"] = errorMessage
		} else {
			fmt.Println("Usuario creado correctamente")
		}
	}
	utils.RenderTemplate(w, "users/new", context)
}

func Login(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if _, err := models.Login(username, password); err != nil {
			context["Error"] = err.Error()
		} else {
			fmt.Println("Login correcto")
		}
	}
	utils.RenderTemplate(w, "users/login", context)
}
