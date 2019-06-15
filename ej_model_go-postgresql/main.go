package main

import (
	"log"
	"time"

	db "./db"
	pg "github.com/go-pg/pg"
)

func main() {
	log.Printf("Listo\n")
	pgDb := db.Conectar()
	SaveProducto(pgDb)
}

//SaveProducto insercion basica en BD
func SaveProducto(dbRef *pg.DB) {
	/*nuevoProducto := &db.Producto{
		Nombre:      "Producto p2",
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
	}*/

	nuevoProducto1 := &db.Producto{
		Nombre:      "Producto p3",
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

	nuevoProducto2 := &db.Producto{
		Nombre:      "Producto p4",
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

	//nuevoProducto.Save(dbRef)
	//nuevoProducto.SaveAndReturn(dbRef)
	totalProds := []*db.Producto{nuevoProducto1, nuevoProducto2}
	nuevoProducto1.SaveMultiple(dbRef, totalProds)
}
