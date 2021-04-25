package main

import (
	"database/sql"
	"log"
	"os"

	// pq es la biblioteca que nos permite conectarnos
	// a postgres con databases/sql.
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	// Obtener la URL de conexión de postgres. Almacenada en una variable de entorno.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Abrir un valor de base de datos. Especificar el controlador postgres para databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sql.Open() no establece ninguna conexión a la base de datos.
	// Simplemente prepara el valor de conexión de la base de datos para su uso posterior.
	// Para asegurar que la base de datos esté disponible y accesible, db.Ping().
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
