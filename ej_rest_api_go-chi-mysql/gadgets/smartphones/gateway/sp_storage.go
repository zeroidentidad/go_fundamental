package gateway

import (
	"github.com/zeroidentidad/rest-chi-mysql/gadgets/smartphones/models"
	"github.com/zeroidentidad/rest-chi-mysql/internal/database"
	"github.com/zeroidentidad/rest-chi-mysql/internal/logs"
)

type SmartphoneStorageGateway interface {
	create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneStorage struct {
	*database.MySqlClient
}

func (s *SmartphoneStorage) create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("Could not create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, os) 
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.OS)

	if err != nil {
		logs.Log().Error("Could not execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("Could not fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		ID:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		OS:            cmd.OS,
	}, nil
}
