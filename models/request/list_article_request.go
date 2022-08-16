package request

import "time"

type ListArticleRequest struct {
	Id    *int32     `form:"id"`
	Start *time.Time `form:"start"`
	End   *time.Time `form:"end"`
	PageRequest
}

func (m *ListArticleRequest) GetNeedSearch() interface{} {
	return *m
}
