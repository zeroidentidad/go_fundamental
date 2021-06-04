package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type OrderService interface {
	Orders(int) (*[]dto.ResponseOrder, int64, *errs.AppError)
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
