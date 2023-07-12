package register

import (
	mail "golang/pkg/mail/gomail"
	registerDto "golang/src/register/dto"
	userDto "golang/src/user/dto"
	userService "golang/src/user/service"
)

type RegisterService interface {
	Register(userDto userDto.CreateUserRequest) error
	VerifyEmail(registerDto registerDto.VerifyEmail) error
}

type RegisterServiceImpl struct {
	userService userService.UserService
	mail        mail.SmtpMail
}

// VerifyEmail implements RegisterService.
func (registerService *RegisterServiceImpl) VerifyEmail(registerDto registerDto.VerifyEmail) error {
	user, err := registerService.userService.FindByEmail(registerDto.EMAIL)

	if err != nil {
		return err
	}

	if user.CodeVerified != registerDto.VERIFICATION_CODE {
		return err
	}

	registerService.userService.Verify(int(user.Id))

	return nil
}

// Register implements RegisterService.
func (registerService *RegisterServiceImpl) Register(userDto userDto.CreateUserRequest) error {

	// create user
	user, err := registerService.userService.Create(userDto)

	// check if user have any error
	if err != nil {
		return err
	}

	// create email dto
	email := registerDto.CreateEmailVerification{
		SUBJECT:           "Welcome",
		EMAIL:             user.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go registerService.mail.SendEmailWelcome(user.Email, email.SUBJECT)

	go registerService.mail.SendVerificationEmail(user.Email, email.VERIFICATION_CODE, email.SUBJECT)

	return nil
}

func NewRegisterService(
	userService userService.UserService,
	mail mail.SmtpMail,
) RegisterService {
	return &RegisterServiceImpl{userService, mail}
}

// type RegisterServiceImpl struct {
// 	userService userService.UserService
// 	mail        mail.SmtpMail
// }

// // register implements RegisterService.
// func (registerService *RegisterServiceImpl) Register(userDto userDto.CreateUserRequestBody) error {
// 	user, err := registerService.userService.Create(userDto)

// 	if err != nil {
// 		return err
// 	}

// 	email := registerDto.CreateEmailVerification{
// 		SUBJECT:           "Welcome to Gobook!",
// 		EMAIL:             user.Email,
// 		VERIFICATION_CODE: "",
// 	}

// 	go registerService.mail.SmtpSendVerificationEmail(user.Email, email)

// 	return nil
// }

// func NewRegisterService(userService userService.UserService, mail mail.SmtpMail) RegisterService {
// 	return &RegisterServiceImpl{userService, mail}
// }
