package domain

type Cliente struct {
	ID              string `db:"cliente_id"`
	Nombre          string
	Ciudad          string
	CodigoPostal    string `db:"codigo_postal"`
	FechaNacimiento string `db:"fecha_nacimiento"`
	Estatus         string
}

type StorageCliente interface {
	FindAll() ([]Cliente, error)
	ById(string) (*Cliente, error)
}
