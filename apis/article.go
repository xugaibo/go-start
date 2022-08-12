package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/api"
	"go-start/db"
	"go-start/models/request"
	"go-start/models/response"
)

type Article struct {
	api.Api
	dao db.ArticleDao
}

func (a Article) List(c *gin.Context) {
	defer a.ErrorHandler()
	a.MakeContext(c)

	param := request.ListArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	r, count := a.dao.List(&param)

	page := response.Page(&r, count, param.PageRequest)

	a.Success(page)
}

func (a Article) Create(c *gin.Context) {
	defer a.ErrorHandler()
	a.MakeContext(c)

	param := request.CreateArticleRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	a.Success(a.dao.Create(param))
}

func (a Article) Delete(c *gin.Context) {
	defer a.ErrorHandler()
	a.MakeContext(c)

	param := request.IdRequest{}
	err := c.ShouldBindUri(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	a.dao.Delete(param.Id)

	a.Ok()
}

func (a Article) Update(c *gin.Context) {
	defer a.ErrorHandler()
	a.MakeContext(c)

	param := request.UpdateArticleRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	a.dao.Update(param)

	a.Ok()
}
