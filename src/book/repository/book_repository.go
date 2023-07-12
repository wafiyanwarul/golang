package book

import (
	bookEntity "golang/src/book/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() []bookEntity.Book
	FindById(id int) (*bookEntity.Book, error)
	Create(book bookEntity.Book) (*bookEntity.Book, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

// Create implements BookRepository.
func (bookRepository *BookRepositoryImpl) Create(book bookEntity.Book) (*bookEntity.Book, error) {

	dataBook := bookRepository.db.Create(&book)

	if dataBook.Error != nil {
		return nil, dataBook.Error
	}

	return &book, nil
}

// FindAll implements BookRepository.
func (bookRepository *BookRepositoryImpl) FindAll() []bookEntity.Book {
	var books []bookEntity.Book

	bookRepository.db.Scopes().Preload("Author").Preload("Publisher").Find(&books)

	return books
}

// FindById implements BookRepository.
func (bookRepository *BookRepositoryImpl) FindById(id int) (*bookEntity.Book, error) {
	var book bookEntity.Book

	dataBook := bookRepository.db.Preload("Author").Preload("Publisher").First(&book, id)

	if dataBook.Error != nil {
		return nil, dataBook.Error
	}

	return &book, nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{db}
}
