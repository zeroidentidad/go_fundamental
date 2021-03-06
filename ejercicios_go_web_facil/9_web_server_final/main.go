package main

import (
	"log"
	"net/http"

	"./config"
	"./handlers"
	v1 "./handlers/api/v1"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	//models.SetDefaultUser()

	//rutas templates html
	mux.HandleFunc("/", handlers.Index).Methods("GET")
	mux.HandleFunc("/users/new", handlers.NewUser).Methods("GET", "POST")
	mux.HandleFunc("/users/login", handlers.Login).Methods("GET", "POST")
	mux.HandleFunc("/users/logout", handlers.Logout).Methods("GET")

	//rutas con middlewares
	editHandler := handlers.Authentication(handlers.UpdateUser)
	editHandler = handlers.MiddlewareTwo(editHandler)
	mux.Handle("/users/edit", editHandler).Methods("GET")

	//endpoints api
	mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", v1.PostUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.PutUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")

	//permitir archivos estaticos
	assets := http.FileServer(http.Dir("assets"))
	statics := http.StripPrefix("/assets/", assets)
	mux.PathPrefix("/assets/").Handler(statics)

	log.Println("Server en puerto:", config.ServerPort())
	server := &http.Server{
		Addr:    config.UrlServer(),
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
