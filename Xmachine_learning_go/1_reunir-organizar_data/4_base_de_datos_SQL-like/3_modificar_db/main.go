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

	// Update algunos registros.
	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// Mostrar cuantas filas fueron actualizadas
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Mostrar numeros de filas en salida estandar
	log.Printf("affected = %d\n", rowCount)
}
