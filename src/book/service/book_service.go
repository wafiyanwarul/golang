package book

import (
	bookDto "golang/src/book/dto"
	bookEntity "golang/src/book/entity"
	bookRepository "golang/src/book/repository"
)

type BookService interface {
	FindAll() []bookEntity.Book
	FindById(id int) (*bookEntity.Book, error)
	Create(bookDto bookDto.CreateBookRequest) (*bookEntity.Book, error)
}

type BookServiceImpl struct {
	bookRepository bookRepository.BookRepository
}

// Create implements BookService.
func (bookService *BookServiceImpl) Create(bookDto bookDto.CreateBookRequest) (*bookEntity.Book, error) {
	var book bookEntity.Book

	book.Title = bookDto.Title
	book.Genre = bookDto.Genre
	book.Stock = bookDto.Stock
	book.AuthorID = bookDto.AuthorID
	book.PublisherID = bookDto.PublisherID

	dataBook, err := bookService.bookRepository.Create(book)

	if err != nil {
		return nil, err
	}

	return dataBook, nil

}

// FindAll implements BookService.
func (bookService *BookServiceImpl) FindAll() []bookEntity.Book {
	return bookService.bookRepository.FindAll()
}

// FindById implements BookService.
func (bookService *BookServiceImpl) FindById(id int) (*bookEntity.Book, error) {
	return bookService.bookRepository.FindById(id)
}

func NewBookService(bookRepository bookRepository.BookRepository) BookService {
	return &BookServiceImpl{bookRepository}
}
