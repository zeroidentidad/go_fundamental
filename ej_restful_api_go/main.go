package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Libro structura de datos base
type Libro struct {
	ID     int    `json:id`
	Titulo string `json:titulo`
	Autor  string `json:autor`
	Anio   string `json:anio`
}

var libros []Libro

func main() {
	router := mux.NewRouter()

	libros = append(libros, Libro{ID: 1, Titulo: "Titulo de prueba 1", Autor: "Anonimo", Anio: "2019"}, Libro{ID: 2, Titulo: "Titulo de prueba 2", Autor: "Anonimo", Anio: "2019"}, Libro{ID: 3, Titulo: "Titulo de prueba 3", Autor: "Anonimo", Anio: "2019"})

	router.HandleFunc("/libros", getLibros).Methods("GET")
	router.HandleFunc("/libros/{id}", getLibro).Methods("GET")
	router.HandleFunc("/libros", addLibro).Methods("POST")
	router.HandleFunc("/libros", updateLibro).Methods("PUT")
	router.HandleFunc("/libros/{id}", removeLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getLibros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(libros)
}

func getLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, libro := range libros {
		if libro.ID == id {
			json.NewEncoder(w).Encode(&libro)
		}
	}
}

func addLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Agregar un libro")
}

func updateLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Actualizar un libro")
}

func removeLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Eliminar un libro")
}
