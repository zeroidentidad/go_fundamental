package main

import (
	"log"
	"time"

	db "./db"
	pg "github.com/go-pg/pg"
)

func main() {
	log.Printf("Listo\n")
	pg_db := db.Conectar()
	SaveProducto(pg_db)
}

func SaveProducto(dbRef *pg.DB) {
	nuevoProducto := &db.Producto{
		Nombre:      "Producto p1",
		Descripcion: "Descripcion producto",
		Imagen:      "/path/imagen",
		Precio:      50.5,
		Caracteristicas: struct {
			Nombre      string
			Descripcion string
			Importe     int
		}{
			Nombre:      "D1",
			Descripcion: "Desc",
			Importe:     3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	nuevoProducto.Save(dbRef)
}
