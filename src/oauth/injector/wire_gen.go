// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oauth

import (
	"golang/src/oauth/delivery/http"
	oauth2 "golang/src/oauth/service"
	"golang/src/user/repository"
	user2 "golang/src/user/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeService(db *gorm.DB) *oauth.OauthHandler {
	userRepository := user.NewUserRepository(db)
	userService := user2.NewUserService(userRepository)
	oauthService := oauth2.NewOauthService(userService)
	oauthHandler := oauth.NewOauthHandler(oauthService)
	return oauthHandler
}
