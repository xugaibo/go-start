package response

import (
	"time"
)

type ListArticleResponse struct {
	Id         int32
	Title      string
	Content    string
	CreateTime time.Time
}
