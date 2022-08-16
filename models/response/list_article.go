package response

import (
	"time"
)

type ListArticle struct {
	Id         int32
	Title      string
	Content    string
	CreateTime time.Time
}
