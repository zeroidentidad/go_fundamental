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
	Valid   bool   `json:"valid"`
}

func HolaJSON(w http.ResponseWriter, r *http.Request) {
	response := CreateResponse("Ke onda es JSON", true)
	json.NewEncoder(w).Encode(response)
}

func CreateResponse(mensaje string, valid bool) Response {
	return Response{
		mensaje,
		valid,
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

	response := Response{}
	if UserExist(username) {
		//no permitir otro ingreso
		response.Mensaje = "No es valido"
		response.Valid = false
	} else {
		//permitir ingreso
		response.Mensaje = "Es valido"
		response.Valid = true
	}
	json.NewEncoder(w).Encode(response)
}

func UserExist(username string) bool {
	Users.RLock()
	defer Users.RUnlock()

	if _, ok := Users.m[username]; ok {
		return true
	}
	return false
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
