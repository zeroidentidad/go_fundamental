package domain

type Cliente struct {
	ID              string
	Nombre          string
	Ciudad          string
	CodigoPostal    string
	FechaNacimiento string
	Estatus         string
}

type StorageCliente interface {
	FindAll() ([]Cliente, error)
}
