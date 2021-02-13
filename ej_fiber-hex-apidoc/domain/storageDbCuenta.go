package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/zeroidentidad/fiber-hex-apidoc/errors"
	"github.com/zeroidentidad/fiber-hex-apidoc/logger"
)

type StorageDbCuenta struct {
	client *sqlx.DB
}

func (d StorageDbCuenta) Save(c Cuenta) (*Cuenta, *errors.AppError) {
	insertSql := "INSERT INTO cuentas(cliente_id, fecha_apertura, tipo_cuenta, cantidad, estatus) values(?,?,?,?,?)"

	result, err := d.client.Exec(insertSql, c.ClienteID, c.FechaApertura, c.TipoCuenta, c.Cantidad, c.Estatus)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	c.ID = strconv.FormatInt(id, 10)
	return &c, nil
}

func NewStorageDbCuenta(dbClient *sqlx.DB) StorageDbCuenta {
	return StorageDbCuenta{dbClient}
}
