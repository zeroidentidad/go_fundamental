package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "database/sql"

	"./models"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var libros []*models.Libro

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("DB_URL"))

	logFatal(err)

	log.Println(pgURL)

	router := mux.NewRouter()

	/*libros = append(libros, &models.Libro{ID: 1, Titulo: "Titulo de prueba 1", Autor: "Anonimo", Anio: 2019}, &models.Libro{ID: 2, Titulo: "Titulo de prueba 2", Autor: "Anonimo", Anio: 2019}, &models.Libro{ID: 3, Titulo: "Titulo de prueba 3", Autor: "Anonimo", Anio: 2019})*/

	router.HandleFunc("/libros", getLibros).Methods("GET")
	router.HandleFunc("/libros/{id}", getLibro).Methods("GET")
	router.HandleFunc("/libros", addLibro).Methods("POST")
	router.HandleFunc("/libros", updateLibro).Methods("PUT")
	router.HandleFunc("/libros/{id}", removeLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	var libro *models.Libro

	_ = json.NewDecoder(r.Body).Decode(&libro)

	libros = append(libros, libro)

	json.NewEncoder(w).Encode(libros)
}

func updateLibro(w http.ResponseWriter, r *http.Request) {
	var libro *models.Libro

	json.NewDecoder(r.Body).Decode(&libro)

	for i, item := range libros {
		if item.ID == libro.ID {
			libros[i] = libro
		}
	}

	json.NewEncoder(w).Encode(libros)
}

func removeLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range libros {
		if item.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(libros)
}
