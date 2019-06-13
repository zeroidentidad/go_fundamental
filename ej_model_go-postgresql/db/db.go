package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

//Conectar a BD postgres
func Conectar() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "x1234567",
		Addr:     "localhost:5432",
		Database: "model_go_pg",
	}
	db := pg.Connect(opts)
	if db == nil {
		log.Printf("Fallo al conectar\n")
		os.Exit(100)
	}
	log.Printf("Conectado\n %v", db)

	CreateProductoTable(db)

	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error al cerrar conexion, motivo: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Coneccion cerrada\n")
}
