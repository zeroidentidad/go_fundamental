package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/utils"
)

// GET /logout
// cerrar sesión del usuario
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning(err, "No se pudo obtener la cookie")
		session := db.Session{Uuid: cookie.Value}
		_ = session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
