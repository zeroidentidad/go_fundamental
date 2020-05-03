package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := httptransport.NewServer(makeGetProductByIdEndPoint(s), getProductByIdRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductsHandler := httptransport.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := httptransport.NewServer(makeAddProductEndPoint(s), addProductRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := httptransport.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := httptransport.NewServer(makeDeleteProductEndPoint(s), deleteProductRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	return r
}

func getProductByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: id,
	}, nil
}

func getProductsRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func addProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func updateProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getUpdateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func deleteProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return getDeleteProductRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}
