package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/helper"
	"webchat/utils"
)

// GET /signup
// Show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	helper.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}
	user := db.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utils.Danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}
