package domain

import (
	"github.com/zeroidentidad/fiber-hex-apidoc/dto"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
)

type Cuenta struct {
	ID            string `db:"cuenta_id"`
	ClienteID     string `db:"cliente_id"`
	FechaApertura string `db:"fecha_apertura"`
	TipoCuenta    string `db:"tipo_cuenta"`
	Cantidad      float64
	Estatus       string
}

func (c Cuenta) ToDtoResponse() dto.ResponseCuenta {
	return dto.ResponseCuenta{ID: c.ID}
}

type StorageCuenta interface {
	Save(Cuenta) (*Cuenta, *errors.AppError)
}
