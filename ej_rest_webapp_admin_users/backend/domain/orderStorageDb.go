package domain

import (
	"backend/errs"
	"backend/logs"
	"errors"

	"gorm.io/gorm"
)

type OrderStorageDb struct {
	client *gorm.DB
}

func NewOrderStorageDb(dbClient *gorm.DB) OrderStorageDb {
	return OrderStorageDb{dbClient}
}

func (db OrderStorageDb) SelectOrders(page int) (*[]Order, int64, *errs.AppError) {
	var total int64
	db.client.Model(&Order{}).Count(&total)
	offset := (page - 1) * LIMIT_PAG
	orders := make([]Order, 0)
	r := db.client.Preload("OrderItems").Offset(offset).Limit(LIMIT_PAG).Find(&orders)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrEmptySlice) {
			return &orders, total, errs.NewUnexpectedError("Orders not found")
		}
		logs.Error("Error finding orders: " + r.Error.Error())
		return &orders, total, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &orders, total, nil
}

func (db OrderStorageDb) SelectAllOrders() (*[]Order, *errs.AppError) {
	orders := make([]Order, 0)
	r := db.client.Preload("OrderItems").Find(&orders)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrEmptySlice) {
			return &orders, errs.NewUnexpectedError("Orders not found")
		}
		logs.Error("Error finding orders: " + r.Error.Error())
		return &orders, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &orders, nil
}
