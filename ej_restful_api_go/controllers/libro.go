package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../repositories"
	"github.com/gorilla/mux"
)

//Controller exported
type Controller struct{}

var libros []models.Libro

//GetLibros exported
func (c Controller) GetLibros(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var libro models.Libro
		libros = []models.Libro{}

		repoLibro := repositories.RepoLibro{}
		libros = repoLibro.GetLibros(db, libro, libros)

		json.NewEncoder(w).Encode(libros)
	}
}

//GetLibro exported
func (c Controller) GetLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var libro models.Libro
		params := mux.Vars(r)

		libros = []models.Libro{}
		repoLibro := repositories.RepoLibro{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		libro = repoLibro.GetLibro(db, libro, id)

		json.NewEncoder(w).Encode(libro)
	}
}

//AddLibro exported
func (c Controller) AddLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var libro models.Libro
		var libroID int

		json.NewDecoder(r.Body).Decode(&libro)

		repoLibro := repositories.RepoLibro{}
		libroID = repoLibro.AddLibro(db, libro)

		json.NewEncoder(w).Encode(libroID)
	}
}

//UpdateLibro exported
func (c Controller) UpdateLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var libro models.Libro

		json.NewDecoder(r.Body).Decode(&libro)

		repoLibro := repositories.RepoLibro{}
		rowsUpdated := repoLibro.UpdateLibro(db, libro)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

//RemoveLibro exported
func (c Controller) RemoveLibro(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		repoLibro := repositories.RepoLibro{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := repoLibro.RemoveLibro(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
