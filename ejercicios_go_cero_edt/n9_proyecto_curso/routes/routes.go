package routes

import "github.com/gorilla/mux"

// InitRoutes es la central que inicia las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)

	return router
}
