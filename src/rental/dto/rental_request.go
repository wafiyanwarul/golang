package rental

import "mime/multipart"

type CreateRentalRequest struct {
	BookID int64 `json:"book_id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

type UpdateRentalRequest struct {
	BookID int64  `json:"book_id" binding:"required"`
	UserID int64  `json:"user_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type ReturnRentalRequest struct {
	Image *multipart.FileHeader `form:"image"`
}
