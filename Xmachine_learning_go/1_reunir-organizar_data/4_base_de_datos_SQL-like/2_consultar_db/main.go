package main

import (
	"database/sql"
	"fmt"
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

	// Query database.
	rows, err := db.Query(`
		SELECT 
			sepal_length as sLength, 
			sepal_width as sWidth, 
			petal_length as pLength, 
			petal_width as pWidth 
		FROM iris
		WHERE species = $1`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterar sobre las filas, enviando los resultados a la salida estándar.
	for rows.Next() {

		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	// Verificar errores después de iterar sobre las filas.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
