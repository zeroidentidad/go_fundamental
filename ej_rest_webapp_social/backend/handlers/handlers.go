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
	router.HandleFunc("/login", middlewares.CheckConnectionDB(routes.Login)).Methods(http.MethodPost)
	router.HandleFunc("/verperfil", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ViewProfile))).Methods(http.MethodGet)
	router.HandleFunc("/modificarperfil", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ModifyProfile))).Methods(http.MethodPut)
	router.HandleFunc("/tweet", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.SendTweet))).Methods(http.MethodPost)
	router.HandleFunc("/vertweets", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ViewTweets))).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = portString
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
