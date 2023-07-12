package oauth

import (
	"errors"
	oauthDto "golang/src/oauth/dto"
	userService "golang/src/user/service"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

)

type OauthService interface {
	Login(oauthDto.LoginRequest) (*oauthDto.LoginResponse, error)
}

type OauthServiceImpl struct {
	userService userService.UserService
}

// Login Implements OauthService.
func (oauthService *OauthServiceImpl) Login(request oauthDto.LoginRequest) (*oauthDto.LoginResponse, error) {
	user, err := oauthService.userService.FindByEmail(request.Email)

	if err != nil {
		return nil, errors.New("invalid email")
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if bcryptErr != nil {
		return nil, errors.New("invalid password")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   expirationTime.Unix(),
	})

	tokenString, signedErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if signedErr != nil {
		return nil, signedErr
	}

	response := oauthDto.LoginResponse{
		Token: tokenString,
		User:  *user,
	}
	return &response, nil
}

// Login implements OauthService.

func NewOauthService(userService userService.UserService) OauthService {
	return &OauthServiceImpl{userService}
}