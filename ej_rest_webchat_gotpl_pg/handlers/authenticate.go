package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/utils"
)

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}

	user, err := db.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		utils.Danger(err, "Cannot find user")
	}

	if user.Password == db.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utils.Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}
