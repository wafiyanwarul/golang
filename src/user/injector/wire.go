//go:build wireinject
// +build wireinject

package user

import (
	userHandler "golang/src/user/delivery/http"
	userRepository "golang/src/user/repository"
	userService "golang/src/user/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *userHandler.UserHandler {
	wire.Build(
		userRepository.NewUserRepository,
		userService.NewUserService,
		userHandler.NewUserHandler,
	)

	return &userHandler.UserHandler{}
}
