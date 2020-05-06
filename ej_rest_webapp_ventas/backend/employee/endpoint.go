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

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getBestEmployeeRequest struct{}

func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndponit := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getEmployeesRequest)

		employees, err := s.GetEmployees(&request)
		helper.Catch(err)

		return employees, nil
	}

	return getEmployeesEndponit
}

func makeGetEmployeeByIdEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIdEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getEmployeeByIDRequest)

		employee, err := s.GetEmployeeById(&request)
		helper.Catch(err)

		return employee, nil
	}

	return getEmployeeByIdEndPoint
}

func makeGetBestEmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		employee, err := s.GetBestEmployee()
		helper.Catch(err)

		return employee, nil
	}

	return getBestEmployeeEndPoint
}
