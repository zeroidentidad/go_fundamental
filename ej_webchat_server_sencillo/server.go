package main

import (
	"log"
	"net/http"

	"encoding/json"
	"sync"

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

func Html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/index.html")
}

type User struct {
	UserName string
}

var Users = struct {
	m map[string]User
	sync.RWMutex
}{m: make(map[string]User)}

func ValidarUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
}

func UserExist(username string) bool {

}

func main() {

	cssHandle := http.FileServer(http.Dir("./frontend/css/"))
	jsHandle := http.FileServer(http.Dir("./frontend/js/"))

	mux := mux.NewRouter()
	mux.HandleFunc("/", Hola).Methods("GET")
	mux.HandleFunc("/json", HolaJSON).Methods("GET")
	mux.HandleFunc("/html", Html).Methods("GET")
	mux.HandleFunc("/validar", ValidarUser).Methods("POST")

	http.Handle("/", mux) //se traducen a las rutas de mux
	http.Handle("/css/", http.StripPrefix("/css/", cssHandle))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandle))

	log.Println("Server ejecutandose")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
