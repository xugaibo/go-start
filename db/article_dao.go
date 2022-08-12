package db

import (
	"errors"
	"go-start/bizcode"
	"go-start/bizerror"
	"go-start/context"
	"go-start/db/models"
	"go-start/enums"
	"go-start/models/request"
	"gorm.io/gorm"
)

type ArticleDao struct {
}

func (d ArticleDao) List(req *request.ListArticleRequest) ([]models.Article, int64) {
	var result []models.Article

	query := context.Db.Model(models.Article{})
	if req.Id > 0 {
		query.Where("id = ?", req.Id)
	}
	query.Where("is_delete = ? ", enums.No.Code())

	var count int64
	query.Count(&count)
	err := query.Limit(req.PageSize).Offset(req.Offset()).Order("id").Find(&result).Error
	if err != nil {
		panic(err)
	}

	return result, count
}

func (d ArticleDao) Create(param request.CreateArticleRequest) uint {
	article := models.Article{}
	article.Content = param.Content
	article.Title = param.Title
	err := context.Db.Create(&article).Error
	if err != nil {
		panic(err)
	}
	return article.Id
}

func (d ArticleDao) Delete(id uint) {
	err := context.Db.Model(models.Article{}).Where("id = ? and is_delete = ? ", id, enums.No.Code()).Update("is_delete", enums.Yes.Code()).Error
	if err != nil {
		panic(err)
	}
}

func (d ArticleDao) Update(param request.UpdateArticleRequest) {
	one := d.FindOne(param.Id)
	if one == nil {
		panic(bizerror.Biz(bizcode.DataNotfound))
	}
	article := models.Article{}
	article.Content = param.Content
	article.Title = param.Title
	article.Id = param.Id
	err := context.Db.Model(&article).Updates(&article).Error
	if err != nil {
		panic(err)
	}
}

func (d ArticleDao) FindOne(id uint) *models.Article {
	article := models.Article{}

	err := context.Db.Where("id = ? and is_delete = ?", id, enums.No.Code()).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return &article
}
