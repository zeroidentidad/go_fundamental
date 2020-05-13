package order

import "github.com/zeroidentidad/backend/helper"

type Service interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(params *addOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetOrderById(param *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(param)
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(params)
	helper.Catch(err)

	totalOrders, err := s.repo.GetTotalOrders(params)
	helper.Catch(err)

	return &OrderList{
		Data:         orders,
		TotalRecords: totalOrders,
	}, nil
}

func (s *service) InsertOrder(params *addOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderId = orderId
		_, err := s.repo.InsertOrderDetail(&detail)
		helper.Catch(err)
	}

	return orderId, nil
}
