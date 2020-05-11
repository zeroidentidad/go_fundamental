package customer

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

	getCustomersHandler := httptransport.NewServer(makeGetCustomersEndPoint(s), getCustomersRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getCustomersHandler)

	return r
}

func getCustomersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getCustomersRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
