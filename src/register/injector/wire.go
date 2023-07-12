//go:build wireinject
// +build wireinject

package register

import (
	registerHandler "golang/src/register/delivery/http"
	registerService "golang/src/register/service"
	userRepository "golang/src/user/repository"
	userService "golang/src/user/service"
	mail "golang/pkg/mail/gomail"

	"github.com/google/wire"
	"gorm.io/gorm"

)

func InitializeService(db *gorm.DB) *registerHandler.RegisterHandler {
	wire.Build(
		registerHandler.NewRegisterHandler,
		userRepository.NewUserRepository,
		registerService.NewRegisterService,
		userService.NewUserService,
		mail.NewSmtpMail,
	)

	return &registerHandler.RegisterHandler{}
}
