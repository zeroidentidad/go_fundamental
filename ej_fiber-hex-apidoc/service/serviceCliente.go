package service

import "github.com/zeroidentidad/fiber-hex-apidoc/domain"

type ServiceCliente interface {
	GetAll() ([]domain.Cliente, error)
	GetById(string) (*domain.Cliente, error)
}

type DefaultServiceCliente struct {
	repo domain.StorageCliente
}

func (s DefaultServiceCliente) GetAll() ([]domain.Cliente, error) {
	return s.repo.FindAll()
}

func (s DefaultServiceCliente) GetById(id string) (*domain.Cliente, error) {
	return s.repo.ById(id)
}

func NewServiceCliente(repo domain.StorageCliente) DefaultServiceCliente {
	return DefaultServiceCliente{
		repo,
	}
}
