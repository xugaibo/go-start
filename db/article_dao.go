package db

import (
	"go-start/context"
	"go-start/db/models"
	"go-start/models/request"
)

type ArticleDao struct {
	models.Article
}

func (d ArticleDao) List(req *request.ListArticleRequest) ([]models.Article, int64) {
	var result []models.Article

	query := context.Db.Model(d.Article)
	if req.Id > 0 {
		query.Where("id = ?", req.Id)
	}

	var count int64
	query.Count(&count)
	query.Limit(req.PageSize).Offset(req.Offset()).Order("id").Find(&result)
	return result, count
}
