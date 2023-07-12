//go:build wireinject
// +build wireinject

package publisher

import (
	publisherHandler "golang/src/publisher/delivery/http"
	publisherRepository "golang/src/publisher/repository"
	publisherService "golang/src/publisher/service"

	"github.com/google/wire"
	"gorm.io/gorm"

)

func InitializeService(db *gorm.DB) *publisherHandler.PublisherHandler {
	wire.Build(
		publisherRepository.NewPublisherRepository,
		publisherService.NewPublisherService,
		publisherHandler.NewPublisherHandler,
	)

	return &publisherHandler.PublisherHandler{}
}
