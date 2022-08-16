package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/bos"
	"go-start/core"
	"go-start/db/user"
	"go-start/models/request"
)

type Token struct {
	core.Api
}

func (t Token) Create(c *gin.Context) {
	defer t.ErrorHandler()
	t.Init(c)

	param := request.CreateToken{}
	err := c.BindJSON(&param)
	if err != nil {
		t.ClientError(err)
		return
	}

	dao := user.NewDao(t.Api)
	userFindByName := dao.GetByUserName(param.UserName)
	bo := bos.UserBo{UserName: param.UserName, Password: param.Password, User: userFindByName}
	m := map[string]string{
		"token": bo.Login(),
	}

	t.Success(m)
}
