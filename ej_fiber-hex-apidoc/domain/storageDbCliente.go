package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type StorageDbCliente struct {
	client *sql.DB
}

func (d StorageDbCliente) FindAll() ([]Cliente, error) {

	findAllSql := "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Query error:", err.Error())
		return nil, err
	}

	clientes := make([]Cliente, 0)
	for rows.Next() {
		c := Cliente{}
		err := rows.Scan(&c.ID, &c.Nombre, &c.Ciudad, &c.CodigoPostal, &c.FechaNacimiento, &c.Estatus)
		if err != nil {
			log.Println("Scan results error:", err.Error())
			return nil, err
		}

		clientes = append(clientes, c)
	}

	return clientes, nil
}

func NewStorageDbCliente() StorageDbCliente {
	client, err := sql.Open("mysql", "root:passx123@tcp(localhost:3306)/banco")
	if err != nil {
		panic(err)
	}

	return StorageDbCliente{
		client,
	}
}
