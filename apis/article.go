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

	param := request.ListArticle{}
	err := c.ShouldBind(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	r, count := dao.List(&param)

	page := response.NewPage(&r, count, param.Page)

	a.Success(page)
}

func (a Article) Create(c *gin.Context) {
	defer a.ErrorHandler()
	a.Init(c)

	param := request.CreateArticle{}
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

	param := request.Id{}
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

	param := request.UpdateArticle{}
	err := c.BindJSON(&param)
	if err != nil {
		a.ClientError(err)
		return
	}

	dao := article.NewDao(a.Api)
	dao.Update(param)

	a.Ok()
}
