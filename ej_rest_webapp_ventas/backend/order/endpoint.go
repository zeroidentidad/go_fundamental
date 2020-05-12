package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getOrderByIdRequest struct {
	orderId int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
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
