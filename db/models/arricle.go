package models

import (
	"errors"
	"go-start/core/bizcode"
	"go-start/core/bizerror"
	"go-start/core/context"
	"go-start/enums"
	"go-start/models/request"
	"gorm.io/gorm"
)

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Base
}

func (d Article) List(req *request.ListArticleRequest) ([]Article, int64) {
	var result []Article

	query := context.Db.Model(Article{})
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

func (d Article) Create(param request.CreateArticleRequest) uint {
	article := Article{}
	article.Content = param.Content
	article.Title = param.Title
	err := context.Db.Create(&article).Error
	if err != nil {
		panic(err)
	}
	return article.Id
}

func (d Article) Delete(id uint) {
	err := context.Db.Model(Article{}).Where("id = ? and is_delete = ? ", id, enums.No.Code()).Update("is_delete", enums.Yes.Code()).Error
	if err != nil {
		panic(err)
	}
}

func (d Article) Update(param request.UpdateArticleRequest) {
	one := d.FindOne(param.Id)
	if one == nil {
		panic(bizerror.Biz(bizcode.DataNotfound))
	}
	article := Article{}
	article.Content = param.Content
	article.Title = param.Title
	article.Id = param.Id
	err := context.Db.Model(&article).Updates(&article).Error
	if err != nil {
		panic(err)
	}
}

func (d Article) FindOne(id uint) *Article {
	article := Article{}

	err := context.Db.Where("id = ? and is_delete = ?", id, enums.No.Code()).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return &article
}
