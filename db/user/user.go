package user

import (
	"go-start/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"userId"`
	UserName  string `json:"userName"`
	UserPhone string `json:"userPhone"`
	Password  string `json:"password"`
	db.Base
}

const (
	PassWordCost = 12 //密码加密难度
)

func (u *User) SetPassword(password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(bytes)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
