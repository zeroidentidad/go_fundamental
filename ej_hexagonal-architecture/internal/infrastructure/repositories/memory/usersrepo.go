package memory

import "github.com/zeroidentidad/hex_arq/internal/core/domain"

type userRepository struct {
}

func NewUsersRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Get() domain.User {
	return domain.User{}
}
