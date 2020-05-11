package customer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	getGetCustomersEndponit := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getCustomersRequest)
		customers, err := s.GetCustomers(&request)
		helper.Catch(err)

		return customers, nil
	}

	return getGetCustomersEndponit
}
