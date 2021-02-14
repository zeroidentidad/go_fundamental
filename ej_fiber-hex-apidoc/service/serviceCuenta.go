package service

import (
	"time"

	"github.com/zeroidentidad/fiber-hex-apidoc/domain"
	"github.com/zeroidentidad/fiber-hex-apidoc/dto"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
)

type ServiceCuenta interface {
	PostNew(dto.RequestCuenta) (*dto.ResponseCuenta, *errors.AppError)
}

type DefaultServiceCuenta struct {
	repo domain.StorageCuenta
}

func (s DefaultServiceCuenta) PostNew(req dto.RequestCuenta) (*dto.ResponseCuenta, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	c := domain.Cuenta{
		ID:            "",
		ClienteID:     req.ClienteID,
		FechaApertura: time.Now().Format("2006-01-02 15:04:05"),
		TipoCuenta:    req.TipoCuenta,
		Cantidad:      req.Cantidad,
		Estatus:       "1",
	}

	cuenta, err := s.repo.Save(c)
	if err != nil {
		return nil, err
	}
	response := cuenta.ToDtoResponse()

	return &response, nil
}

func NewServiceCuenta(repo domain.StorageCuenta) DefaultServiceCuenta {
	return DefaultServiceCuenta{
		repo,
	}
}
