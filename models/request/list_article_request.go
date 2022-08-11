package request

type ListArticleRequest struct {
	Id int32 `form:"id"`
	PageRequest
}

func (m *ListArticleRequest) GetNeedSearch() interface{} {
	return *m
}
