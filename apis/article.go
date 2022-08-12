package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/api"
	"go-start/db"
	"go-start/models/request"
	"go-start/models/response"
	"net/http"
)

type Article struct {
	api.Api
}

func (a Article) List(c *gin.Context) {
	defer a.ErrorHandler()
	a.MakeContext(c)

	param := request.ListArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		a.ClientError("can not parse param")
		return
	}

	dao := db.ArticleDao{}
	r, count := dao.List(&param)

	page := response.Page(&r, count, param.PageRequest)
	c.JSON(http.StatusOK, response.Ok(page))
}

func (a Article) Create(c *gin.Context) {
	r := true
	c.JSON(http.StatusOK, r)
}

func (a Article) Delete(c *gin.Context) {
	r := true
	c.JSON(http.StatusOK, r)
}

func (a Article) Update(c *gin.Context) {
	r := true
	c.JSON(http.StatusOK, r)
}
