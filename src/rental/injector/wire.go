//go:build wireinject
// +build wireinject

package rental

import (
	fileupload "golang/pkg/fileupload/cloudinary"
	rentalHandler "golang/src/rental/delivery/http"
	rentalRepository "golang/src/rental/repository"
	rentalService "golang/src/rental/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *rentalHandler.RentalHandler {
	wire.Build(
		rentalRepository.NewRentalRepository,
		rentalService.NewRentalService,
		rentalHandler.NewRentalHandler,
		fileupload.NewFileUpload,
	)

	return &rentalHandler.RentalHandler{}
}
