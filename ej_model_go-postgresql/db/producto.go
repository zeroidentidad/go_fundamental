package db

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*Producto name migration: producto */
type Producto struct {
	RefPointer      int      `sql:"-"`
	tableName       struct{} `sql:"productos"`
	ID              int      `sql:"id,pk"`
	Nombre          string   `sql:"nombre,unique"`
	Descripcion     string   `sql:"descripcion"`
	Imagen          string   `sql:"imagen"`
	Precio          float64  `sql:"precio,type:real"`
	Caracteristicas struct {
		Nombre      string
		Descripcion string
		Importe     int
	} `sql:"caracteristicas,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

//CreateProductoTable usar modelo para crear tabla
func CreateProductoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&Producto{}, opts)
	if createErr != nil {
		log.Printf("Error creando tabla. Error: %v\n", createErr)

		return createErr
	}
	log.Printf("Tabla creada")
	return nil
}

//Save para inserta la estructura de datos del producto
func (pi *Producto) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error insertando elemmento a BD, %v\n", insertErr)
		return insertErr
	}
	log.Printf("Producto %s insertado\n", pi.Nombre)

	return nil
}
