package main

import (
	"log"
	"net/http"

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
	router.HandleFunc("/libros", getLibros).Methods("GET")
	router.HandleFunc("/libros/{id}", getLibro).Methods("GET")
	router.HandleFunc("/libros", addLibro).Methods("POST")
	router.HandleFunc("/libros", updateLibro).Methods("PUT")
	router.HandleFunc("/libros/{id}", removeLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getLibros(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtener todos los libros")
}

func getLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtener un libro")
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
