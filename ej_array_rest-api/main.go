package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" //bajar con go get ó con go-dep
)

//Persona data model
type Persona struct {
	ID        int        `json:"id,omitempty"`
	Nombre    string     `json:"nombre,omitempty"`
	Apellido  string     `json:"apellido,omitempty"`
	Direccion *Direccion `json:"direccion,omitempty"`
}

// Direccion data model
type Direccion struct {
	Ciudad string `json:"ciudad,omitempty"`
	Estado string `json:"estado,omitempty"`
}

var people []Persona

func main() {
	router := mux.NewRouter()

	// adding example data
	people = append(people, Persona{ID: 1, Nombre: "Nombre X", Apellido: "Xorres", Direccion: &Direccion{Ciudad: "Rancho", Estado: "Tabasco"}})
	people = append(people, Persona{ID: 2, Nombre: "Nombre X2", Apellido: "Xonzalez"})
	people = append(people, Persona{ID: 2, Nombre: "Nombre X3", Apellido: "Xerrer", Direccion: &Direccion{Ciudad: "Puerto", Estado: "Veracruz"}})

	// endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

// GetPeopleEndpoint toda la coleccion
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint un elemento
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Persona{})
}

// CreatePersonEndpoint nuevo elemento
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	var person Persona
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = id
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}

// DeletePersonEndpoint remover un elemento
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
