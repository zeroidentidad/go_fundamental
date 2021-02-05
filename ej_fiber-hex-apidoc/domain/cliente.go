package domain

import "github.com/zeroidentidad/fiber-hex-apidoc/errors"

type Cliente struct {
	ID              string `db:"cliente_id"`
	Nombre          string
	Ciudad          string
	CodigoPostal    string `db:"codigo_postal"`
	FechaNacimiento string `db:"fecha_nacimiento"`
	Estatus         string
}

type StorageCliente interface {
	FindAll(string) ([]Cliente, *errors.AppError)
	ById(string) (*Cliente, *errors.AppError)
}
