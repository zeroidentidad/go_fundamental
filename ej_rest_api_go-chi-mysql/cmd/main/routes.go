package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/rest-chi-mysql/gadgets/smartphones/web"
	reviews "github.com/zeroidentidad/rest-chi-mysql/reviews/web"
)

func Routes(
	sph *web.CreateSmartphoneHandler,
	reviewHandler *reviews.ReviewHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log cada solicitud http
		middleware.Recoverer, // recuperar si ocurre panic
	)

	mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)
	mux.Post("/reviews", reviewHandler.AddReviewHandler)

	return mux
}

// test handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "zero")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
