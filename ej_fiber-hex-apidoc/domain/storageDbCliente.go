package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
)

type StorageDbCliente struct {
	client *sql.DB
}

func (d StorageDbCliente) FindAll(estatus string) ([]Cliente, *errors.AppError) {
	var findAllSql string
	if estatus == "" {
		findAllSql = "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes"
	} else {
		findAllSql = "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes WHERE estatus = " + estatus
	}

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		return nil, errors.NewNotFoundError("Customers not found")
	}

	clientes := make([]Cliente, 0)
	for rows.Next() {
		c := Cliente{}
		err := rows.Scan(&c.ID, &c.Nombre, &c.Ciudad, &c.CodigoPostal, &c.FechaNacimiento, &c.Estatus)
		if err != nil {
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}

		clientes = append(clientes, c)
	}

	return clientes, nil
}

func (d StorageDbCliente) ById(id string) (*Cliente, *errors.AppError) {
	findByIdSql := "SELECT cliente_id, nombre, ciudad, codigo_postal, fecha_nacimiento, estatus FROM clientes WHERE cliente_id = ?"

	row := d.client.QueryRow(findByIdSql, id)

	var c Cliente
	err := row.Scan(&c.ID, &c.Nombre, &c.Ciudad, &c.CodigoPostal, &c.FechaNacimiento, &c.Estatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
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
