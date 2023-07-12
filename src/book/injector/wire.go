//go:build wireinject
// +build wireinject

package book

import (
	bookHandler "golang/src/book/delivery/http"
	bookRepository "golang/src/book/repository"
	bookService "golang/src/book/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *bookHandler.BookHandler {
	wire.Build(
		bookRepository.NewBookRepository,
		bookService.NewBookService,
		bookHandler.NewBookHandler,
	)

	return &bookHandler.BookHandler{}
}
