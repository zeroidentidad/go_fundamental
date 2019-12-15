package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// ruta agregada
	r.POST("/user", createUser)
	// ruta agregada mas parametro
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<a href="/user/9872309847">Ir a: http://localhost:8080/user/9872309847</a>
			</body>
			</html>
		`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	// Marshal en JSON
	ujson, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", ujson)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// literal compuesta - tipo y llaves
	u := models.User{}

	// encode/decode para enviar/recibir JSON a/desde una secuencia de bytes
	json.NewDecoder(r.Body).Decode(&u)

	// cambiar ID
	u.ID = "007"

	// marshal/unmarshal por tener JSON asignado a una variable
	ujson, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", ujson)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: escribir código para eliminar usuario
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Escribir código para eliminar usuario\n")
}

/*
Start your server

Enter this at the terminal:

curl http://localhost:8080/user/1

curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user

-X abreviatura de --request - Especifica un método de solicitud personalizado para usar cuando se comunica con el servidor HTTP.

-H abreviatura de --header

-d abreviatura de --data

curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
*/
