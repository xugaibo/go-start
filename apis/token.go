package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/bos"
	"go-start/core"
	"go-start/db/models"
	"go-start/models/request"
)

type Token struct {
	core.Api
	dao models.User
}

func (u Token) Create(c *gin.Context) {
	defer u.ErrorHandler()
	u.MakeContext(c)

	param := request.CreateTokenRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		u.ClientError(err)
		return
	}

	userFindByName := u.dao.GetByUserName(param.UserName)
	bo := bos.UserBo{UserName: param.UserName, Password: param.Password, UserFindByName: userFindByName}
	m := map[string]string{
		"token": bo.Login(),
	}

	u.Success(m)
}
