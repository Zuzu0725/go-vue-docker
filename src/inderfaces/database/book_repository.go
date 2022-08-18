package database

import "www/src/domain"

type BookRepository struct {
	SqlHandler
}

func (repo *BookRepository) FindById(id int) (book domain.Book, err error) {
	if err = repo.Find(&book, id).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) FindAll() (books domain.Books, err error) {
	if err = repo.Find(&books).Error; err != nil {
		return
	}
	return
}

func (repo *BookRepository) Store(b domain.Book) (book domain.Book, err error) {
	if err = repo.Create(&b).Error; err != nil {
		return
	}
	book = b
	return
}

func (repo *BookRepository) Update(b domain.Book) (book domain.Book, err error) {
	if err = repo.Save(&b).Error; err != nil {
		return
	}
	book = b
	return
}

func (repo *BookRepository) DeleteById(book domain.Book) (err error) {
	if err = repo.Delete(&book).Error; err != nil {
		return
	}
	return
}
