package service

import (
	"github.com/zeroidentidad/fiber-hex-apidoc/domain"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
)

type ServiceCliente interface {
	GetAll(string) ([]domain.Cliente, *errors.AppError)
	GetById(string) (*domain.Cliente, *errors.AppError)
}

type DefaultServiceCliente struct {
	repo domain.StorageCliente
}

func (s DefaultServiceCliente) GetAll(estatus string) ([]domain.Cliente, *errors.AppError) {
	return s.repo.FindAll(estatus)
}

func (s DefaultServiceCliente) GetById(id string) (*domain.Cliente, *errors.AppError) {
	return s.repo.ById(id)
}

func NewServiceCliente(repo domain.StorageCliente) DefaultServiceCliente {
	return DefaultServiceCliente{
		repo,
	}
}
