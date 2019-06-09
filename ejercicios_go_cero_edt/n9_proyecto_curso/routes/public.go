package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SetPublicRouter expone en el web server los archivos estaticos
func SetPublicRouter(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
