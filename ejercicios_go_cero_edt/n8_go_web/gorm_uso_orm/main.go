package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Producto contiene los datos de un artículo
type Producto struct {
	gorm.Model   //ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("mysql", "remoto:x1234567@/go_edtest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Error de conexion a BD")
	}
	defer db.Close()
	fmt.Println("Se conecto a la BD")

	/*db.CreateTable(&Producto{})
	fmt.Println("Se creo la tabla producto en la BD")*/

	//db.Create(&Producto{CodigoBarras: "010101234455", Precio: 7234})

	var p Producto
	db.First(&p, 2)
	fmt.Println("Precio del producto consultado:", p.Precio)

	p.Precio = 14500
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es", p.Precio)
}
