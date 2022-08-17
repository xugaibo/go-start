package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/bos"
	"go-start/core"
	"go-start/core/jwtutil/refreshtoken"
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
	token, refreshToken := bo.Login()
	m := map[string]string{
		"token":        token,
		"refreshToken": refreshToken,
	}

	t.Success(m)
}

func (t Token) Update(c *gin.Context) {
	defer t.ErrorHandler()
	t.Init(c)

	param := request.UpdateToken{}
	err := c.BindJSON(&param)
	if err != nil {
		t.ClientError(err)
		return
	}

	m := map[string]string{
		"token": refreshtoken.Refresh(param.RefreshToken),
	}

	t.Success(m)
}
