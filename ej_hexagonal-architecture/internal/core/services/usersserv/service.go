package usersserv

import (
	"github.com/zeroidentidad/hex_arq/internal/core/domain"
	"github.com/zeroidentidad/hex_arq/internal/core/ports"
)

type service struct {
	repo ports.UserRepository
}

func NewService(r ports.UserRepository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) Get() domain.User {
	return domain.User{}
}
