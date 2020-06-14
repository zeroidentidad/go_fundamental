package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zeroidentidad/twittor/config"
	"github.com/zeroidentidad/twittor/middlewares"
	"github.com/zeroidentidad/twittor/routes"
)

var confObj = *config.GetConfiguration()
var portString = fmt.Sprintf("%s", confObj.HttpPort)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckConnectionDB(routes.Registry)).Methods(http.MethodPost)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = portString
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
