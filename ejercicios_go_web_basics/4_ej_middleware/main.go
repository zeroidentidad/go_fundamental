package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {

	// Middleware stack simple
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":8080")

}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Iniciando sesión aqui ...")
	if r.URL.Query().Get("password") == "secretx" {
		next(rw, r)
	} else {
		http.Error(rw, "No autorizado", 401)
	}
	log.Println("Iniciar sesión nuevamente...")
}
