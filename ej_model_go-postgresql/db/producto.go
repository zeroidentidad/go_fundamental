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

//SaveAndReturn para inserta y retornar valores usados
func (pi *Producto) SaveAndReturn(db *pg.DB) (*Producto, error) {
	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()

	if insertErr != nil {
		log.Printf("Error insertando elemmento a BD, %v\n", insertErr)
		return nil, insertErr
	}
	log.Printf("Nuevo Producto insertado\n")
	log.Printf("Los valores recibidos fueron: %v\n", InsertResult.RowsReturned())

	return pi, nil
}

//SaveMultiple registro multiples elementos
func (pi *Producto) SaveMultiple(db *pg.DB, items []*Producto) error {
	_, insertErr := db.Model(items[0], items[1]).Insert()

	if insertErr != nil {
		log.Printf("Error insertando registros, %v\n", insertErr)
		return insertErr
	}
	log.Printf("Registros insertados\n")

	return nil
}

//DeleteItem elimina usando el nombre del model
func (pi *Producto) DeleteItem(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("nombre = ?nombre").Delete()

	if deleteErr != nil {
		log.Printf("Error eliminando registro, %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Registro eliminado\n", pi.Nombre)

	return nil
}

//UpdatePrice actualiza el precio de un producto
func (pi *Producto) UpdatePrice(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("precio = ?precio").Where("id = ?id").Update()

	if updateErr != nil {
		log.Printf("Error actualizando registro, %v\n", updateErr)
		return updateErr
	}
	log.Printf("Precio del registro actualizado ID: %d\n", pi.ID)

	return nil
}

//GetByID obtener estructura de datos por el id proporcionado
func (pi *Producto) GetByID(db *pg.DB) error {

	//getErr := db.Select(pi)

	//getErr := db.Model(pi).Where("id = ?0", pi.ID).Select()

	/*getErr := db.Model(pi).Column("nombre", "descripcion").
	Where("id = ?0", pi.ID).
	WhereOr("nombre = ?0", pi.Nombre).Select()*/

	var items []Producto
	getErr := db.Model(&items).Column("nombre", "descripcion").
		Where("id = ?0", pi.ID).
		WhereOr("id = ?0", 1).
		Offset(0).
		Limit(2).
		Order("id desc").
		Select()

	if getErr != nil {
		log.Printf("Error obteniendo valores por el id, %v\n", getErr)
		return getErr
	}

	log.Printf("Consulta correcta para: %v\n", items) //*pi

	return nil
}
