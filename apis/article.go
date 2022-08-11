package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/db"
	"go-start/models/request"
	"go-start/models/response"
	"net/http"
)

type Article struct {
	db.ArticleDao
}

func (a Article) List(c *gin.Context) {
	param := request.ListArticleRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "can not parse param")
	}
	r, count := a.ArticleDao.List(&param)
	c.JSON(http.StatusOK, response.Of(&r, count, param.PageRequest))
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
