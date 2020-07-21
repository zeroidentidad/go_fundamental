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

	// User crud data
	router.HandleFunc("/registro", middlewares.CheckConnectionDB(routes.Registry)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlewares.CheckConnectionDB(routes.Login)).Methods(http.MethodPost)
	router.HandleFunc("/verperfil", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ViewProfile))).Methods(http.MethodGet)
	router.HandleFunc("/modificarperfil", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ModifyProfile))).Methods(http.MethodPut)
	// Tweets crud data
	router.HandleFunc("/tweet", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.SendTweet))).Methods(http.MethodPost)
	router.HandleFunc("/vertweets", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ViewTweets))).Methods(http.MethodGet)
	router.HandleFunc("/eliminartweet", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.RemoveTweet))).Methods(http.MethodDelete)
	// Avatars image
	router.HandleFunc("/subiravatar", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.UploadAvatar))).Methods(http.MethodPost)
	router.HandleFunc("/obteneravatar", middlewares.CheckConnectionDB(routes.FindAvatar)).Methods(http.MethodGet)
	// Banners image
	router.HandleFunc("/subirbanner", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.UploadBanner))).Methods(http.MethodPost)
	router.HandleFunc("/obtenerbanner", middlewares.CheckConnectionDB(routes.FindBanner)).Methods(http.MethodGet)
	// Relaciones users
	router.HandleFunc("/relacionar", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.SetRelation))).Methods(http.MethodPost)
	router.HandleFunc("/desrelacionar", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.RemoveRelation))).Methods(http.MethodDelete)
	router.HandleFunc("/checkrelacion", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.GetRelation))).Methods(http.MethodGet)
	router.HandleFunc("/listausuarios", middlewares.CheckConnectionDB(middlewares.ValidateJWT(routes.ListUsers))).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = portString
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
