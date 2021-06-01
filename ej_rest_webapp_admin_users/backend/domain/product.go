package domain

import (
	"backend/dto"
	"backend/errs"
)

type Product struct {
	ID          uint    `gorm:"column:id;primaryKey"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Image       string  `gorm:"column:image"`
	Price       float64 `gorm:"column:price"`
}

type ProductStorage interface {
	InsertProduct(Product) (*Product, *errs.AppError)
	SelectProduct(Product) (*Product, *errs.AppError)
	SelectProducts(int) (*[]Product, int64, *errs.AppError)
	UpdateProduct(Product) (*Product, *errs.AppError)
	DeleteProduct(Product) *errs.AppError
}

func NewProduct(id uint, title, desc, img string, price float64) Product {
	return Product{
		ID:          id,
		Title:       title,
		Description: desc,
		Image:       img,
		Price:       price,
	}
}

func (p Product) ToDto() dto.ResponseProduct {
	return dto.ResponseProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
	}
}
