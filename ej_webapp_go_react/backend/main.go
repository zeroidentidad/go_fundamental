package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtgoreact",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	router.Handle("agregar canal", agregarCanal)
	router.Handle("subscribir canal", subscribirCanal)
	router.Handle("desubscribir canal", desubscribirCanal)

	router.Handle("editar usuario", editarUsuario)
	router.Handle("subscribir usuario", subscribirUsuario)
	router.Handle("desubscribir usuario", desubscribirUsuario)

	router.Handle("agregar mensaje", agregarMensaje)
	router.Handle("subscribir mensaje", subscribirMensaje)
	router.Handle("desubscribir mensaje", desubscribirMensaje)
	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
