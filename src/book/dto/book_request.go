package book

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Stock       int64  `json:"stock" binding:"required"`
	PublisherID int64  `json:"publisher_id" binding:"required"`
	AuthorID    int64  `json:"author_id" binding:"required"`
}
