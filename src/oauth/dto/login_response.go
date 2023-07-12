package oauth

import (
	userEntity "golang/src/user/entity"
)

type LoginResponse struct {
	Token string          `json:"token"`
	User  userEntity.User `json:"user"`
}
