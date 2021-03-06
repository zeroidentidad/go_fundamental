package product

import "github.com/zeroidentidad/backend/helper"

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int64, error)
	DeleteProduct(params *getDeleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopList, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)
	helper.Catch(err)

	totalProducts, err := s.repo.GetTotalProducts()
	helper.Catch(err)

	return &ProductList{Data: products, TotalRecords: totalProducts}, err
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *getUpdateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *getDeleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}

func (s *service) GetBestSellers() (*ProductTopList, error) {
	productsTop, err := s.repo.GetBestSellers()
	helper.Catch(err)

	totalVentas, err := s.repo.GetTotalVentas()
	helper.Catch(err)

	return &ProductTopList{Data: productsTop, TotalVentas: totalVentas}, err
}
