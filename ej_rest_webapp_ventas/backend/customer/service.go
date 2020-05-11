package customer

import "github.com/zeroidentidad/backend/helper"

type Service interface {
	GetCustomers(params *getCustomersRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetCustomers(params *getCustomersRequest) (*CustomerList, error) {
	customers, err := s.repo.GetCustomers(params)
	helper.Catch(err)

	totalCustomers, err := s.repo.GetTotalCustomers()
	helper.Catch(err)

	return &CustomerList{
		Data:         customers,
		TotalRecords: totalCustomers,
	}, nil
}
