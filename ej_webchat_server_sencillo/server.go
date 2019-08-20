package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ke pedo"))
}

type Response struct {
	Mensaje string `json:"mensaje"`
	Nombre  string `json:"nombre"`
}

func HolaJSON(w http.ResponseWriter, r *http.Request) {
	resonse := CreateResponse()
	json.NewEncoder(w).Encode(resonse)
}

func CreateResponse() Response {
	return Response{
		"Ke onda es JSON",
		"Jesus",
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Hola).Methods("GET")
	mux.HandleFunc("/json", HolaJSON).Methods("GET")

	http.Handle("/", mux) //se traducen a las rutas de mux

	log.Println("Server ejecutandose")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
