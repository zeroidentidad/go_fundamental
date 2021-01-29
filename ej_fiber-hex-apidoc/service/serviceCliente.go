package service

import "github.com/zeroidentidad/fiber-hex-apidoc/domain"

type ServiceCliente interface {
	GetAll() ([]domain.Cliente, error)
}

type DefaultServiceCliente struct {
	repo domain.StorageCliente
}

func (s DefaultServiceCliente) GetAll() ([]domain.Cliente, error) {
	return s.repo.FindAll()
}

func NewServiceCliente(repo domain.StorageCliente) DefaultServiceCliente {
	return DefaultServiceCliente{
		repo,
	}
}
