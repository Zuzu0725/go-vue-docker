package usecase

import "www/src/domain"

type BookRepository interface {
	FindById(id int) (domain.Book, error)
	FindAll() (domain.Books, error)
	Store(domain.Book) (domain.Book, error)
	Update(domain.Book) (domain.Book, error)
	DeleteById(domain.Book) error
}
