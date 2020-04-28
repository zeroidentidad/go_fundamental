package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/backend/database"
	"github.com/zeroidentidad/backend/product"
)

var dbConn *sql.DB

func main() {

	dbConn = database.InitDB()
	defer dbConn.Close()

	var productRepository = product.NewRepository(dbConn)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
