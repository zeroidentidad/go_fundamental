package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./commons"
	"./migration"
	"./routes"
	"github.com/urfave/negroni"
)

// inicio por defecto: ./run_app --port 8080
// inicio en blanco: ./run_app --port 8081 --migrate yes
func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera migración a BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración.")
	}

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución programa")
}
