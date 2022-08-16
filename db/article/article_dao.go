package article

import (
	"errors"
	"go-start/core"
	"go-start/core/bizcode"
	"go-start/core/bizerror"
	"go-start/core/enums"
	"go-start/db"
	"go-start/models/request"
	"gorm.io/gorm"
)

type Dao struct {
	db.BaseDao
}

func NewDao(api core.Api) *Dao {
	dao := Dao{}
	api.MakeBaseDao(&dao.BaseDao)
	return &dao
}

func (d *Dao) List(req *request.ListArticleRequest) ([]Article, int64) {
	var result []Article

	query := d.Db.Model(Article{})
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

func (d *Dao) Create(param request.CreateArticleRequest) uint {
	m := Article{}
	m.Content = param.Content
	m.Title = param.Title
	err := d.Db.Create(&m).Error
	if err != nil {
		panic(err)
	}
	return m.Id
}

func (d *Dao) Delete(id uint) {
	err := d.Db.Model(Article{}).Where("id = ? and is_delete = ? ", id, enums.No.Code()).Update("is_delete", enums.Yes.Code()).Error
	if err != nil {
		panic(err)
	}
}

func (d *Dao) Update(param request.UpdateArticleRequest) {
	one := d.FindOne(param.Id)
	if one == nil {
		panic(bizerror.Biz(bizcode.DataNotfound))
	}
	m := Article{}
	m.Content = param.Content
	m.Title = param.Title
	err := d.Db.Where("id = ?", param.Id).Updates(&m).Error
	if err != nil {
		panic(err)
	}
}

func (d *Dao) FindOne(id uint) *Article {
	m := Article{}

	err := d.Db.Where("id = ? and is_delete = ?", id, enums.No.Code()).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}

	return &m
}
