package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./conexion"
	"./estructuras"
	"github.com/gorilla/mux"
)

func main() {
	conexion.InitDB()
	defer conexion.ConnClose()

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/nuevo", NewUser).Methods("POST")

	puerto := ":3000"
	log.Println("Servidor iniciado, puerto", puerto)
	log.Fatal(http.ListenAndServe(puerto, r))
}

//GetUser func
func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idUsuario := vars["id"]

	usuario := conexion.GetUser(idUsuario)
	//usuario := estructuras.Usuario{1, "Jesus", "Ferrer", 26}

	status := "success"
	var mensaje = "ok"
	if usuario.Id <= 0 {
		status = "error"
		mensaje = "Usuario no encontrado."
	}
	respuesta := estructuras.Respuesta{status, usuario, mensaje}
	json.NewEncoder(w).Encode(respuesta)
}

//NewUser func
func NewUser(w http.ResponseWriter, r *http.Request) {
	usuario := GetUserRequest(r)

	respuesta := estructuras.Respuesta{"success", conexion.CreateUser(usuario), "ok"}
	json.NewEncoder(w).Encode(respuesta)
}

func GetUserRequest(r *http.Request) estructuras.Usuario {
	var usuario estructuras.Usuario
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&usuario)
	if err != nil {
		log.Fatal(err)
	}

	return usuario
}
