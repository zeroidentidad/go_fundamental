package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetUser).Methods("GET")

	puerto := ":3000"
	log.Println("Servidor iniciado, puerto", puerto)
	log.Fatal(http.ListenAndServe(puerto, r))
}

//GetUser func
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("epale"))
}
