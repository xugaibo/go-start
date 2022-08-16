package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/bos"
	"go-start/core"
	"go-start/db/user"
	"go-start/models/request"
)

type User struct {
	core.Api
}

func (u User) Create(c *gin.Context) {
	defer u.ErrorHandler()
	u.Init(c)

	param := request.CreateUserRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		u.ClientError(err)
		return
	}

	dao := user.NewDao(u.Api)
	userFindByName := dao.GetByUserName(param.UserName)
	bo := bos.UserBo{UserName: param.UserName, UserPhone: param.UserPhone, Password: param.Password, UserFindByName: userFindByName}

	u.Success(dao.Create(bo.NewUser()))
}
