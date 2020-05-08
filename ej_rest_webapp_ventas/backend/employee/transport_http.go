package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zeroidentidad/backend/helper"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEmployeesHandler := httptransport.NewServer(makeGetEmployeesEndPoint(s), getEmployeesRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeByIDHandler := httptransport.NewServer(makeGetEmployeeByIdEndPoint(s), getEmployeeByIDRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIDHandler)

	getBestEmployeeHandler := httptransport.NewServer(makeGetBestEmployeeEndPoint(s), getBestEmployeeRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeHandler)

	addEmployeeHandler := httptransport.NewServer(makeInsertEmployeeEndPoint(s), addEmployeeRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addEmployeeHandler)

	updateEmployeeHandler := httptransport.NewServer(makeUpdateEmployeeEndPoint(s), updateEmployeeRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateEmployeeHandler)

	deleteEmployeeHandler := httptransport.NewServer(makeDeleteEmployeeEndPoint(s), deleteEmployeeRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteEmployeeHandler)

	return r
}

func getEmployeesRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func getEmployeeByIDRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}

func getBestEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}

func addEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func updateEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeeRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func deleteEmployeeRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return deleteEmployeeRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}
