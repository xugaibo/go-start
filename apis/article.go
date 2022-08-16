package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/core"
	"go-start/db/article"
	"go-start/models/request"
	"go-start/models/response"
)

type Article struct {
	core.Api
}

func (a Article) List(c *gin.Context) {
	defer a.ErrorHandler()
	a.Init(c)

	param := request.ListArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	r, count := dao.List(&param)

	page := response.Page(&r, count, param.PageRequest)

	a.Success(page)
}

func (a Article) Create(c *gin.Context) {
	defer a.ErrorHandler()
	a.Init(c)

	param := request.CreateArticleRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	a.Success(dao.Create(param))
}

func (a Article) Delete(c *gin.Context) {
	defer a.ErrorHandler()
	a.Init(c)

	param := request.IdRequest{}
	err := c.ShouldBindUri(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	dao.Delete(param.Id)

	a.Ok()
}

func (a Article) Update(c *gin.Context) {
	defer a.ErrorHandler()
	a.Init(c)

	param := request.UpdateArticleRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	dao.Update(param)

	a.Ok()
}
