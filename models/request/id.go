package request

type Id struct {
	Id uint `uri:"id" binding:"required"`
}
