package request

import "time"

type ListArticle struct {
	Id    *int32     `form:"id"`
	Start *time.Time `form:"start"`
	End   *time.Time `form:"end"`
	Page
}

func (m *ListArticle) GetNeedSearch() interface{} {
	return *m
}
