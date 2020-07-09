package request

type CreateLinkRequest struct {
	Url string `json:"url" binding:"required"`
}
