package order

type Service interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetOrderById(param *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(param)
}
