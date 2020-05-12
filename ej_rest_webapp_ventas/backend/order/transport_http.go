package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zeroidentidad/backend/helper"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdHandler := httptransport.NewServer(makeGetOrderByIdEndPoint(s), getOrderByIdRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	getOrdersHandler := httptransport.NewServer(makeGetOrdersEndPoint(s), getOrdersRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	return r
}

func getOrderByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)

	return getOrderByIdRequest{
		orderId: orderId,
	}, nil
}

func getOrdersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
