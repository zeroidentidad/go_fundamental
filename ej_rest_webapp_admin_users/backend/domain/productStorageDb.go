package domain

import (
	"backend/errs"
	"backend/logs"
	"errors"

	"gorm.io/gorm"
)

type ProductStorageDb struct {
	client *gorm.DB
}

func NewProductStorageDb(dbClient *gorm.DB) ProductStorageDb {
	return ProductStorageDb{dbClient}
}

func (db ProductStorageDb) InsertProduct(p Product) (*Product, *errs.AppError) {
	r := db.client.Create(&p)
	if r.Error != nil {
		logs.Error("Error creating product: " + r.Error.Error())
		return &p, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (db ProductStorageDb) SelectProduct(p Product) (*Product, *errs.AppError) {
	var product Product
	r := db.client.Where("id = ?", p.ID).First(&product)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return &product, errs.NewUnexpectedError("Product not found")
		}
		logs.Error("Error finding product: " + r.Error.Error())
		return &product, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &product, nil
}

func (db ProductStorageDb) SelectProducts(page int) (*[]Product, int64, *errs.AppError) {
	var total int64
	db.client.Model(&Product{}).Count(&total)
	offset := (page - 1) * LIMIT_PAG
	products := make([]Product, 0)
	r := db.client.Offset(offset).Limit(LIMIT_PAG).Find(&products)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrEmptySlice) {
			return &products, total, errs.NewUnexpectedError("Products not found")
		}
		logs.Error("Error finding products: " + r.Error.Error())
		return &products, total, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &products, total, nil
}

func (db ProductStorageDb) UpdateProduct(p Product) (*Product, *errs.AppError) {
	r := db.client.Model(&p).Updates(p)
	if r.Error != nil {
		logs.Error("Error updating product: " + r.Error.Error())
		return &p, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (db ProductStorageDb) DeleteProduct(p Product) *errs.AppError {
	r := db.client.Delete(&p)
	if r.Error != nil {
		logs.Error("Error deleting product: " + r.Error.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}
