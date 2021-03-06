package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zeroidentidad/backend/helper"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type getUpdateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type getDeleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct{}

// @Summary Producto by ID
// @Tags Products
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {object} product.Product "OK"
// @Router /products/{id} [get]
func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getProductByIDRequest)
		product, err := s.GetProductById(&request)

		helper.Catch(err)

		return product, nil
	}

	return getProductByIdEndPoint
}

// @Summary Lista de Productos
// @Tags Products
// @Accept json
// @Produce  json
// @Param request body product.getProductsRequest true "User Data"
// @Success 200 {object} product.ProductList "OK"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getProductsRequest)
		products, err := s.GetProducts(&request)

		helper.Catch(err)

		return products, nil
	}

	return getProductsEndPoint
}

// @Summary Insertar Producto
// @Tags Products
// @Accept json
// @Produce  json
// @Param request body product.getAddProductRequest true "User Data"
// @Success 200 {integer} int "OK"
// @Router /products/ [post]
func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getAddProductRequest)
		productId, err := s.InsertProduct(&request)

		helper.Catch(err)

		return productId, nil
	}

	return addProductEndPoint
}

// @Summary Update Producto
// @Tags Products
// @Accept json
// @Produce  json
// @Param request body product.getUpdateProductRequest true "User Data"
// @Success 200 {integer} int "OK"
// @Router /products/ [put]
func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getUpdateProductRequest)
		productId, err := s.UpdateProduct(&request)

		helper.Catch(err)

		return productId, nil
	}

	return updateProductEndPoint
}

// @Summary Eliminar Producto
// @Tags Products
// @Accept json
// @Produce  json
// @Param id path int true "Producto Id"
// @Success 200 {integer} int "OK"
// @Router /products/{id} [delete]
func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(getDeleteProductRequest)
		rows, err := s.DeleteProduct(&request)

		helper.Catch(err)

		return rows, nil
	}

	return deleteProductEndPoint
}

// @Summary Best Sellers
// @Tags Products
// @Accept json
// @Produce  json
// @Success 200 {object} product.ProductTopList "OK"
// @Router /products/bestSellers [get]
func makeBestSellersEndPoint(s Service) endpoint.Endpoint {
	getBestSellersEndPoint := func(ctx context.Context, req interface{}) (interface{}, error) {
		productsTop, err := s.GetBestSellers()

		helper.Catch(err)

		return productsTop, nil
	}

	return getBestSellersEndPoint
}
