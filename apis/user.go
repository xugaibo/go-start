package apis

import (
	"github.com/gin-gonic/gin"
	"go-start/bos"
	"go-start/core"
	"go-start/db/models"
	"go-start/models/request"
)

type User struct {
	core.Api
	dao models.User
}

func (u User) Create(c *gin.Context) {
	defer u.ErrorHandler()
	u.MakeContext(c)

	param := request.CreateUserRequest{}
	err := c.BindJSON(&param)
	if err != nil {
		u.ClientError(err)
		return
	}

	userFindByName := u.dao.GetByUserName(param.UserName)
	bo := bos.UserBo{UserName: param.UserName, UserPhone: param.UserPhone, Password: param.Password, UserFindByName: userFindByName}

	u.Success(u.dao.Create(bo.NewUser()))
}
