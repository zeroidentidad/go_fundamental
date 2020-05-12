package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getOrderByIdRequest struct {
	orderId int64
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
