package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers"
	"./drivers"
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

	controller := controllers.Controller{}

	router.HandleFunc("/libros", controller.GetLibros(db)).Methods("GET")
	router.HandleFunc("/libros/{id}", controller.GetLibro(db)).Methods("GET")
	router.HandleFunc("/libros", controller.AddLibro(db)).Methods("POST")
	router.HandleFunc("/libros", controller.UpdateLibro(db)).Methods("PUT")
	router.HandleFunc("/libros/{id}", controller.RemoveLibro(db)).Methods("DELETE")

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
