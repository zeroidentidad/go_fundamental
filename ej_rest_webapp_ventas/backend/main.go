package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/backend/database"
	"github.com/zeroidentidad/backend/employee"
	"github.com/zeroidentidad/backend/product"
)

var dbConn *sql.DB

func main() {

	dbConn = database.InitDB()
	defer dbConn.Close()

	var (
		productRepository = product.NewRepository(dbConn)
		productService    product.Service
	)
	productService = product.NewService(productRepository)

	var (
		employeeRepository = employee.NewRepository(dbConn)
		employeeService    employee.Service
	)
	employeeService = employee.NewService(employeeRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))

	http.ListenAndServe(":3000", r)
}

// live-reload bin watcher: github.com/canthefason/go-watcher
