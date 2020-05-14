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

type deleteEmployeeRequest struct {
	EmployeeID string
}

// @Summary Lista de Empleados
// @Tags Employees
// @Accept json
// @Produce  json
// @Param request body employee.getEmployeesRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "OK"
// @Router /employees/paginated [post]
func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndponit := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getEmployeesRequest)

		employees, err := s.GetEmployees(&request)
		helper.Catch(err)

		return employees, nil
	}

	return getEmployeesEndponit
}

// @Summary Empleado by ID
// @Tags Employees
// @Accept json
// @Produce  json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "ok"
// @Router /employees/{id} [get]
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

func makeDeleteEmployeeEndPoint(s Service) endpoint.Endpoint {
	getDeleteEmployeeEndPoint := func(_ context.Context, req interface{}) (interface{}, error) {
		request := req.(deleteEmployeeRequest)
		employee, err := s.DeleteEmployee(&request)
		helper.Catch(err)

		return employee, nil
	}

	return getDeleteEmployeeEndPoint
}
