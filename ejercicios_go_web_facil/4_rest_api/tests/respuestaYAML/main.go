package main

import (
	"log"
	"net/http"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	mux := mux.NewRouter()

	//endpoints
	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	log.Println("Server en puerto:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
