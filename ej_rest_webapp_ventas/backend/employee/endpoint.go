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

type addEmployeeRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type updateEmployeeRequest struct {
	ID            int64
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
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

func makeInsertEmployeeEndPoint(s Service) endpoint.Endpoint {
	getInsertEmployeeEndPoint := func(_ context.Context, req interface{}) (interface{}, error) {
		request := req.(addEmployeeRequest)
		employee, err := s.InsertEmployee(&request)
		helper.Catch(err)

		return employee, nil
	}

	return getInsertEmployeeEndPoint
}

func makeUpdateEmployeeEndPoint(s Service) endpoint.Endpoint {
	getUpdateEmployeeEndPoint := func(_ context.Context, req interface{}) (interface{}, error) {
		request := req.(updateEmployeeRequest)
		employee, err := s.UpdateEmployee(&request)
		helper.Catch(err)

		return employee, nil
	}

	return getUpdateEmployeeEndPoint
}
