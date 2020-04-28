package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
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
