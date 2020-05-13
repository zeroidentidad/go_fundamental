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

// @Summary Lista de Clientes
// @Tags Customers
// @Accept json
// @Produce json
// @Param request body customer.getCustomersRequest true "User Data"
// @Success 200 {object} customer.CustomerList "OK"
// @Router /customers/paginated [post]
func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	getGetCustomersEndponit := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getCustomersRequest)
		customers, err := s.GetCustomers(&request)
		helper.Catch(err)

		return customers, nil
	}

	return getGetCustomersEndponit
}
