package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
	"strconv"
)

type ProductService interface {
	Products(int) (*[]dto.ResponseProduct, int64, *errs.AppError)
	CreateProduct(dto.RequestProduct) (*dto.ResponseProduct, *errs.AppError)
	GetProduct(string) (*dto.ResponseProduct, *errs.AppError)
	UpdateProduct(dto.RequestProduct) (*dto.ResponseProduct, *errs.AppError)
	DeleteProduct(string) *errs.AppError
}

type DefaultProductService struct {
	repo domain.ProductStorage
}

func NewProductService(repo domain.ProductStorage) DefaultProductService {
	return DefaultProductService{
		repo,
	}
}

func (s DefaultProductService) Products(page int) (*[]dto.ResponseProduct, int64, *errs.AppError) {
	res := make([]dto.ResponseProduct, 0)
	products, total, err := s.repo.SelectProducts(page)
	if err != nil {
		return &res, total, err
	}

	for _, p := range *products {
		res = append(res, p.ToDto())
	}

	return &res, total, err
}

func (s DefaultProductService) CreateProduct(req dto.RequestProduct) (res *dto.ResponseProduct, err *errs.AppError) {
	p := domain.NewProduct(req.ID, req.Title, req.Description, req.Image, req.Price)

	product, err := s.repo.InsertProduct(p)
	if err != nil {
		return res, err
	}

	dto := product.ToDto()

	return &dto, nil
}

func (s DefaultProductService) GetProduct(id string) (res *dto.ResponseProduct, err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	p := domain.NewProduct(uint(_id), "", "", "", 0)

	product, err := s.repo.SelectProduct(p)
	if err != nil {
		return res, err
	}

	dto := product.ToDto()

	return &dto, nil
}

func (s DefaultProductService) UpdateProduct(req dto.RequestProduct) (res *dto.ResponseProduct, err *errs.AppError) {
	p := domain.NewProduct(req.ID, req.Title, req.Description, req.Image, req.Price)

	product, err := s.repo.UpdateProduct(p)
	if err != nil {
		return res, err
	}

	dto := product.ToDto()

	return &dto, nil
}

func (s DefaultProductService) DeleteProduct(id string) (err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	p := domain.NewProduct(uint(_id), "", "", "", 0)

	err = s.repo.DeleteProduct(p)
	if err != nil {
		return err
	}

	return nil
}
