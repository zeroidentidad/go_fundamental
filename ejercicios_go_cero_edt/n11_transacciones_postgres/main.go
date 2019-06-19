package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//Factura modelo
type Factura struct {
	ID      int
	Cliente string
	Fecha   time.Time
}

//FacturaDetalle modelo
type FacturaDetalle struct {
	ID         int
	FacturaID  int
	ProductoID int
	Cantidad   int
}

func main() {
	f := Factura{
		Cliente: "Jesus GO",
	}

	fd := FacturaDetalle{
		ProductoID: 1,
		Cantidad:   2,
	}

	// Obtener DB conexion
	db := getConnection()
	defer db.Close()

	// Crear transaccion
	tx, err := db.Begin()
	if err != nil {
		log.Printf("No se pudo crear transaccion. Error: %v", err)
	}

	// Insercion encabezado
	id, err := InsertarFactura(f, tx)
	if err != nil {
		tx.Rollback()
		fmt.Println("Ocurrio un error en encabezado: " + err.Error())
		return
	}

	// Insercion detalle
	fd.FacturaID = id
	err = InsertarFacturaDetalle(fd, tx)
	if err != nil {
		tx.Rollback()
		fmt.Println("Ocurrio un error en detalle: " + err.Error())
		return
	}

	tx.Commit()
	fmt.Println("Transaccion realizada correctamente")
}

//InsertarFactura maestro
func InsertarFactura(f Factura, tx *sql.Tx) (int, error) {
	var id int
	q := "INSERT INTO invoices (client) VALUES ($1) RETURNING id"

	err := tx.QueryRow(q, f.Cliente).Scan(&id)
	if err != nil {
		log.Printf("Error en insercion de factura. Error: %v\n", err)
		return id, err
	}

	return id, nil
}

//InsertarFacturaDetalle detalle
func InsertarFacturaDetalle(fd FacturaDetalle, tx *sql.Tx) error {
	q := `INSERT INTO 
			invoices_details (invoice_id, product_id, amount)
			VALUES ($1, $2, $3)`

	result, err := tx.Exec(q, fd.FacturaID, fd.ProductoID, fd.Cantidad)
	if err != nil {
		log.Printf("Error en insercion de detalle. Error: %v\n", err)
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error al obtener registros afectados. Error: %v\n", err)
		return err
	}

	if i != 1 {
		err = fmt.Errorf("Se esperaba obtener una fila afectada, se obtuvieron: %d", i)
		log.Println(err)
		return err
	}

	return nil
}

func getConnection() *sql.DB {
	con := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", "postgres", "x1234567", "localhost", "5432", "transacciones_go", "disable")

	db, err := sql.Open("postgres", con)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
