package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundos"))
})

var PostRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Es una solicitud Post"))
})

var PathVariableHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("Hola " + name))
})

func main() {
	router := mux.NewRouter()
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods(http.MethodGet)

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods(http.MethodPost)
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, PathVariableHandler)).Methods(http.MethodGet)
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
