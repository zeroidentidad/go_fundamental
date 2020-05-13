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

	addOrderHandler := httptransport.NewServer(makeAddOrderEndPoint(s), addOrderRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addOrderHandler)

	updateOrderHandler := httptransport.NewServer(makeUpdateOrderEndPoint(s), updateOrderRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateOrderHandler)

	return r
}

func getOrderByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)

	return getOrderByIdRequest{
		OrderId: orderId,
	}, nil
}

func getOrdersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func addOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}

func updateOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
