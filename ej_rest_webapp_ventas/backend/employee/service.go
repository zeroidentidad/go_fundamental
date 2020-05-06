package employee

import "github.com/zeroidentidad/backend/helper"

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)

	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)

	return &EmployeeList{
		Data:         employees,
		TotalRecords: totalEmployees,
	}, nil
}

func (s *service) GetEmployeeById(param *getEmployeeByIDRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(param)
}
