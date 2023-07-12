//go:build wireinject
// +build wireinject

package author

import (
	authorHandler "golang/src/author/delivery/http"
	authorRepository "golang/src/author/repository"
	authorService "golang/src/author/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *authorHandler.AuthorHandler {
	wire.Build(
		authorRepository.NewAuthorRepository,
		authorService.NewAuthorService,
		authorHandler.NewAuthorHandler,
	)

	return &authorHandler.AuthorHandler{}
}
