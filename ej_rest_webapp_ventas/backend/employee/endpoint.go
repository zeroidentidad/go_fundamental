package employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndponit := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getEmployeesRequest)

		employees, err := s.GetEmployees(&request)
		helper.Catch(err)

		return employees, nil
	}

	return getEmployeesEndponit
}
