package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Instance string `json:"instance" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UpdateUserRequestBody struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Instance string `json:"instance"`
}
