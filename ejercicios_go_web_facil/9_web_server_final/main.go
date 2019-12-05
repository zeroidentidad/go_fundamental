package main

import (
	"log"
	"net/http"

	"./config"
	v1 "./handlers/api/v1"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	//models.SetDefaultUser()

	//endpoints
	mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", v1.PostUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.PutUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")

	log.Println("Server en puerto:", config.ServerPort())
	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
