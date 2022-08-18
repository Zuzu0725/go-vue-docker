package usecase

import "www/src/domain"

type BookInteractor struct {
	BookRepository BookRepository
}

func (interactor *BookInteractor) BookById(id int) (book domain.Book, err error) {
	book, err = interactor.BookRepository.FindById(id)
	return
}

func (interactor *BookInteractor) Books() (books domain.Books, err error) {
	books, err = interactor.BookRepository.FindAll()
	return
}

func (interactor *BookInteractor) Add(b domain.Book) (book domain.Book, err error) {
	book, err = interactor.BookRepository.Store(b)
	return
}

func (interactor *BookInteractor) Update(b domain.Book) (book domain.Book, err error) {
	book, err = interactor.BookRepository.Update(b)
	return
}

func (interactor *BookInteractor) DeletebyId(b domain.Book) (err error) {
	err = interactor.BookRepository.DeleteById(b)
	return
}
