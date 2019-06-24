package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	//"./models"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

//Libro structura de datos base
type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Anio   int    `json:"anio"`
}

var libros []Libro
var db *sql.DB

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("DB_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

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
	var libro Libro
	libros = []Libro{}

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
	var libro Libro
	params := mux.Vars(r)

	rows := db.QueryRow("select * from libros where id=$1", params["id"])
	err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Anio)
	logFatal(err)

	json.NewEncoder(w).Encode(libro)
}

func addLibro(w http.ResponseWriter, r *http.Request) {
	var libro Libro
	var bookID int

	json.NewDecoder(r.Body).Decode(&libro)

	err := db.QueryRow("insert into libros (titulo, autor, anio) values ($1, $2, $3) returning id;", libro.Titulo, libro.Autor, libro.Anio).Scan(&bookID)

	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateLibro(w http.ResponseWriter, r *http.Request) {
	var libro Libro

	json.NewDecoder(r.Body).Decode(&libro)

	result, err := db.Exec("update libros set titulo=$1, autor=$2, anio=$3 where id=$4 returning id", libro.Titulo, libro.Autor, libro.Anio, libro.ID)

	rowUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowUpdated)
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
