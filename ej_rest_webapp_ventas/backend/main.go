package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zeroidentidad/backend/customer"
	"github.com/zeroidentidad/backend/database"
	_ "github.com/zeroidentidad/backend/docs"
	"github.com/zeroidentidad/backend/employee"
	"github.com/zeroidentidad/backend/order"
	"github.com/zeroidentidad/backend/product"
)

// @title Rest-Go-Chi-Kit API
// @version 1.0
// @description Este es un servidor api ventas de ejemplo.
// @contact.name ZeroIdentidad
// @contact.url https://zeroidentidad.github.io/chat

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

	// docs swag init
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)
}

// opt1: live-reload bin watcher: github.com/canthefason/go-watcher
// opt2: live-reload bin air: github.com/cosmtrek/air
