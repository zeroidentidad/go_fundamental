package handlers

import (
	"net/http"
	"webchat/db"
	"webchat/utils"
)

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning(err, "Failed to get cookie")
		session := db.Session{Uuid: cookie.Value}
		_ = session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
