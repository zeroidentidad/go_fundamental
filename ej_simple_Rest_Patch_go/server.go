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
