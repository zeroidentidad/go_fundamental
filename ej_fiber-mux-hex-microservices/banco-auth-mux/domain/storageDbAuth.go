package domain

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type StorageAuth interface {
	FindBy(username string, password string) (*Login, error)
}

type StorageDbAuth struct {
	client *sqlx.DB
}

func (d StorageDbAuth) FindBy(username, password string) (*Login, error) {
	var login Login
	sqlVerify := `SELECT username, u.cliente_id, rol, group_concat(c.cuenta_id) as numeros_cuenta FROM usuarios u
	LEFT JOIN cuentas c ON c.cliente_id = u.cliente_id WHERE username = ? and password = ?
	GROUP BY c.cliente_id`

	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			return nil, errors.New("Unexpected database error")
		}
	}
	return &login, nil
}

func NewStorageDbAuth(client *sqlx.DB) StorageDbAuth {
	return StorageDbAuth{client}
}
