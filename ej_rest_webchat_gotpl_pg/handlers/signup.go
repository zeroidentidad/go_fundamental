package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/helper"
	"webchat/utils"
)

// GET /signup
// mostrar página de registro
func Signup(writer http.ResponseWriter, request *http.Request) {
	helper.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// crear cuenta de usuario
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "No se puede procesar formulario")
	}
	user := db.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utils.Danger(err, "No se puede crear usuario")
	}
	http.Redirect(writer, request, "/login", 302)
}
