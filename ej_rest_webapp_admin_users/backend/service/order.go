package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type OrderService interface {
	Orders(int) (*[]dto.ResponseOrder, int64, *errs.AppError)
	ExportOrders() (*[]dto.ResponseOrder, *errs.AppError)
	ChartSales() (*[]dto.ResponseSales, *errs.AppError)
}

type DefaultOrderService struct {
	repo domain.OrderStorage
}

func NewOrderService(repo domain.OrderStorage) DefaultOrderService {
	return DefaultOrderService{
		repo,
	}
}

func (s DefaultOrderService) Orders(page int) (*[]dto.ResponseOrder, int64, *errs.AppError) {
	res := make([]dto.ResponseOrder, 0)
	orders, total, err := s.repo.SelectOrders(page)
	if err != nil {
		return &res, total, err
	}

	for _, o := range *orders {
		res = append(res, o.ToDto())
	}

	return &res, total, err
}

func (s DefaultOrderService) ExportOrders() (*[]dto.ResponseOrder, *errs.AppError) {
	res := make([]dto.ResponseOrder, 0)
	orders, err := s.repo.SelectAllOrders()
	if err != nil {
		return &res, err
	}

	for _, o := range *orders {
		res = append(res, o.ToDto())
	}

	return &res, err
}

func (s DefaultOrderService) ChartSales() (*[]dto.ResponseSales, *errs.AppError) {
	res := make([]dto.ResponseSales, 0)
	sales, err := s.repo.SelectSales()
	if err != nil {
		return &res, err
	}

	for _, s := range *sales {
		res = append(res, s.ToDto())
	}

	return &res, err
}
