package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getOrderByIdRequest struct {
	OrderId int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

type addOrderRequest struct {
	ID           int64
	OrderDate    string
	CustomerID   int
	OrderDetails []addOrderDetailsRequest
}

type addOrderDetailsRequest struct {
	ID        int64
	OrderId   int64
	ProductId int64
	Quantity  int64
	UnitPrice float64
}

func makeGetOrderByIdEndPoint(s Service) endpoint.Endpoint {
	getOrderByIdEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getOrderByIdRequest)

		order, err := s.GetOrderById(&request)
		helper.Catch(err)

		return order, nil
	}

	return getOrderByIdEndPoint
}

func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	getOrdersEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getOrdersRequest)

		orders, err := s.GetOrders(&request)
		helper.Catch(err)

		return orders, nil
	}

	return getOrdersEndPoint
}

func makeAddOrderEndPoint(s Service) endpoint.Endpoint {
	addOrderEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(addOrderRequest)

		order, err := s.InsertOrder(&request)
		helper.Catch(err)

		return order, nil
	}

	return addOrderEndPoint
}

func makeUpdateOrderEndPoint(s Service) endpoint.Endpoint {
	updateOrderEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(addOrderRequest)

		order, err := s.UpdateOrder(&request)
		helper.Catch(err)

		return order, nil
	}

	return updateOrderEndPoint
}
