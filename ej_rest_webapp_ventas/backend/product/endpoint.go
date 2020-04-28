package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getProductByIDRequest)
		product, err := s.GetProductById(&request)

		if err != nil {
			panic(err)
		}

		return product, nil
	}

	return getProductByIdEndPoint
}
