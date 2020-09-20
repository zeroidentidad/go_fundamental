package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/utils"
)

// POST /authenticate
// autenticar usuario dado el correo electrónico y la contraseña
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "No se puede procesar formulario")
	}

	user, err := db.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		utils.Danger(err, "No se puede encontrar al usuario")
	}

	if user.Password == db.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utils.Danger(err, "No se puede crear sesión")
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
