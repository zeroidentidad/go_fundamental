// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ZeroIdentidad",
            "url": "https://zeroidentidad.github.io/chat"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Lista de Clientes",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.getCustomersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerList"
                        }
                    }
                }
            }
        },
        "/employees/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Update Empleado",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.updateEmployeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Insertar Empleado",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.addEmployeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/employees/best": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Mejor Empleado",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/employee.BestEmployee"
                        }
                    }
                }
            }
        },
        "/employees/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Lista de Empleados",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.getEmployeesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/employee.EmployeeList"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Empleado by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/employee.Employee"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Eliminar Empleado",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Update Order",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.addOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Insertar Order",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.addOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Lista de Ordenes",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.getOrdersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.OrderList"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Order by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.OrderItem"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Eliminar Orden [cabecera y detalle]",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}/detail/{orderDetailId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Eliminar elemento detalle de Orden",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Detail Id",
                        "name": "orderDetailId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/products/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update Producto",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.getUpdateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Insertar Producto",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.getAddProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/products/bestSellers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Best Sellers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.ProductTopList"
                        }
                    }
                }
            }
        },
        "/products/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Lista de Productos",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.getProductsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.ProductList"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Producto by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Producto Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Eliminar Producto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Producto Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "customer.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessphone": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "customer.CustomerList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Customer"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "customer.getCustomersRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "employee.BestEmployee": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "totalVentas": {
                    "type": "integer"
                }
            }
        },
        "employee.Employee": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "employee.EmployeeList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/employee.Employee"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "employee.addEmployeeRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "employee.getEmployeesRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "employee.updateEmployeeRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "order.OrderDetailItem": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "order_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        },
        "order.OrderItem": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "customer": {
                    "type": "string"
                },
                "customerId": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderDetailItem"
                    }
                },
                "orderDate": {
                    "type": "string"
                },
                "orderId": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "statusId": {
                    "type": "string"
                },
                "statusName": {
                    "type": "string"
                }
            }
        },
        "order.OrderList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderItem"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "order.addOrderDetailsRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "orderId": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "number"
                }
            }
        },
        "order.addOrderRequest": {
            "type": "object",
            "properties": {
                "customerID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "orderDate": {
                    "type": "string"
                },
                "orderDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.addOrderDetailsRequest"
                    }
                }
            }
        },
        "order.getOrdersRequest": {
            "type": "object",
            "properties": {
                "dateFrom": {
                    "type": "object"
                },
                "dateTo": {
                    "type": "object"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "status": {
                    "type": "object"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "listPrice": {
                    "type": "number"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standardCost": {
                    "type": "number"
                }
            }
        },
        "product.ProductList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/product.Product"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "product.ProductTop": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "productName": {
                    "type": "string"
                },
                "vendidos": {
                    "type": "number"
                }
            }
        },
        "product.ProductTopList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/product.ProductTop"
                    }
                },
                "totalVentas": {
                    "type": "number"
                }
            }
        },
        "product.getAddProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "listPrice": {
                    "type": "number"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standardCost": {
                    "type": "number"
                }
            }
        },
        "product.getProductsRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "product.getUpdateProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "listPrice": {
                    "type": "number"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standardCost": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Rest-Go-Chi-Kit API",
	Description: "Este es un servidor api ventas de ejemplo.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
