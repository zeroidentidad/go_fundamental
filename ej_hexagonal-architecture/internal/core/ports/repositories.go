package ports

import "github.com/zeroidentidad/hex_arq/internal/core/domain"

type PersonsRepository interface {
	Get(id int) (domain.Person, error)
	Save(person domain.Person) error
}

type UserRepository interface {
	Get() domain.User
}
