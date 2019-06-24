package main

import (
	"database/sql"
	"encoding/json"
	"go_fundamental/ej_restful_api_go/drivers"
	"log"
	"net/http"

	"./models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var libros []models.Libro
var db *sql.DB

func main() {

	db = drivers.ConectarDB()
	log.Println("running") //pgURL

	router := mux.NewRouter()

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
	var libro models.Libro
	libros = []models.Libro{}

	rows, err := db.Query("select * from libros")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Anio)
		logFatal(err)

		libros = append(libros, libro)
	}

	json.NewEncoder(w).Encode(libros)
}

func getLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	params := mux.Vars(r)

	rows := db.QueryRow("select * from libros where id=$1", params["id"])
	err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Anio)
	logFatal(err)

	json.NewEncoder(w).Encode(libro)
}

func addLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	var bookID int

	json.NewDecoder(r.Body).Decode(&libro)

	err := db.QueryRow("insert into libros (titulo, autor, anio) values ($1, $2, $3) returning id;", libro.Titulo, libro.Autor, libro.Anio).Scan(&bookID)

	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro

	json.NewDecoder(r.Body).Decode(&libro)

	result, err := db.Exec("update libros set titulo=$1, autor=$2, anio=$3 where id=$4 returning id", libro.Titulo, libro.Autor, libro.Anio, libro.ID)

	rowUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowUpdated)
}

func removeLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from libros where id=$1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
