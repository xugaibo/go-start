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

func (d *Dao) List(req *request.ListArticle) ([]Article, int64) {
	var result []Article

	query := d.Db.Debug().Model(Article{})
	query.Where("is_delete = ? ", enums.No.Code())

	if req.Id != nil {
		query.Where("id = ?", req.Id)
	}
	if req.Start != nil {
		query.Where("created_at >= ", req.Start)
	}
	if req.End != nil {
		query.Where("updated_at <=", req.End)
	}

	var count int64
	query.Count(&count)
	err := query.Limit(req.PageSize).Offset(req.Offset()).Order("created_at desc").Find(&result).Error
	if err != nil {
		panic(err)
	}

	return result, count
}

func (d *Dao) Create(param request.CreateArticle) uint {
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

func (d *Dao) Update(param request.UpdateArticle) {
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
