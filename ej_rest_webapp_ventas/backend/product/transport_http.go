package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := httptransport.NewServer(makeGetProductByIdEndPoint(s), getProductByIdRequestDecoder, httptransport.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	return r
}

func getProductByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: id,
	}, nil
}
