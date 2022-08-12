package models

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Base
}

func (Article) TableName() string {
	return "article"
}
