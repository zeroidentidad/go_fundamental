package local

import (
	"../book"
)

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
