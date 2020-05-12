package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/backend/customer"
	"github.com/zeroidentidad/backend/database"
	"github.com/zeroidentidad/backend/employee"
	"github.com/zeroidentidad/backend/order"
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

	var (
		customerRepository = customer.NewRepository(dbConn)
		customerService    customer.Service
	)
	customerService = customer.NewService(customerRepository)

	var (
		orderRepository = order.NewRepository(dbConn)
		orderService    order.Service
	)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))

	http.ListenAndServe(":3000", r)
}

// live-reload bin watcher: github.com/canthefason/go-watcher
