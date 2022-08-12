package request

type UpdateArticleRequest struct {
	Id uint `json:"id" binding:"required"`
	CreateArticleRequest
}
