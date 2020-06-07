package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zeroidentidad/twittor/config"
)

var confObj = *config.GetConfiguration()
var portString = fmt.Sprintf("%s", confObj.HttpPort)

func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = portString
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
