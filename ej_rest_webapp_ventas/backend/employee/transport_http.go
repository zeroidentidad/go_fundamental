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

	return r
}

func getEmployeesRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
