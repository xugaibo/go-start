package request

type Page struct {
	PageSize   int `form:"pageSize"`
	PageNumber int `form:"pageNumber"`
}

func (req Page) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}
