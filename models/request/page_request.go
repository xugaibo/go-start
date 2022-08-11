package request

type PageRequest struct {
	PageSize   int `form:"pageSize"`
	PageNumber int `form:"pageNumber"`
}

func (req PageRequest) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}
