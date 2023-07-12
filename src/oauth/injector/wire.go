//go:build wireinject
// +build wireinject

package oauth

import (
	userRepository "golang/src/user/repository"
	userService "golang/src/user/service"
	oauthHandler "golang/src/oauth/delivery/http"
	oauthService "golang/src/oauth/service"

	"github.com/google/wire"
	"gorm.io/gorm"

)

func InitializeService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthService.NewOauthService,
		oauthHandler.NewOauthHandler,
		userService.NewUserService,
		userRepository.NewUserRepository,
	)

	return &oauthHandler.OauthHandler{}
}
