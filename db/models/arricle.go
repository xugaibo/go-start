package models

type Article struct {
	Id      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (Article) TableName() string {
	return "article"
}
