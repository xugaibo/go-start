package request

type IdRequest struct {
	Id uint `uri:"id" binding:"required"`
}
