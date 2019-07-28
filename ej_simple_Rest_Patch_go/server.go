package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./conexion"
	"github.com/gorilla/mux"
)

type Usuario struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

func main() {
	conexion.InitDB()
	defer conexion.ConnClose()

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")

	puerto := ":3000"
	log.Println("Servidor iniciado, puerto", puerto)
	log.Fatal(http.ListenAndServe(puerto, r))
}

//GetUser func
func GetUser(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	//id_usuario := vars["id"]

	usuario := Usuario{"Jesus", "Ferrer", 26}
	json.NewEncoder(w).Encode(usuario)
}
