package bos

import (
	"go-start/core/bizcode"
	"go-start/core/bizerror"
	"go-start/core/jwtutil"
	"go-start/core/util"
	"go-start/db/models"
)

type UserBo struct {
	UserName       string
	UserPhone      string
	Password       string
	UserFindByName *models.User
}

func (r UserBo) NewUser() *models.User {
	if r.UserFindByName != nil {
		panic(bizerror.Biz(bizcode.UserNameExists))
	}

	if r.UserPhone != "" && !util.VerifyMobile(r.UserPhone) {
		panic(bizerror.Biz(bizcode.PhoneInvalid))
	}

	user := models.User{UserName: r.UserName, UserPhone: r.UserPhone}
	user.SetPassword(r.Password)
	return &user
}

func (r UserBo) Login() string {
	if r.UserFindByName == nil {
		panic(bizerror.Biz(bizcode.UserNotExists))
	}

	if !r.UserFindByName.CheckPassword(r.Password) {
		panic(bizerror.Biz(bizcode.PasswordInvalid))
	}

	token, err := jwtutil.GenerateToken(r.UserFindByName.UserId, r.UserName)
	if err != nil {
		panic(err)
	}
	return token
}
