package publisher

type CreatePublisherRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}
