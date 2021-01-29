package domain

type StorageStubCliente struct {
	clientes []Cliente
}

func (s StorageStubCliente) FindAll() ([]Cliente, error) {
	return s.clientes, nil
}

func NewStorageStubCliente() StorageStubCliente {
	clientes := []Cliente{
		{"1001", "Jesus", "Tabasco", "86690", "1900-01-01", "1"},
		{"1002", "Otro", "Tabasco", "86699", "1901-01-01", "1"},
	}

	return StorageStubCliente{
		clientes,
	}
}
