package article

import (
	"go-start/db"
)

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	db.Base
}
