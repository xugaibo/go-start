package request

type UpdateArticle struct {
	Id uint `json:"id" binding:"required"`
	CreateArticle
}
