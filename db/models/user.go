package models

import (
	"errors"
	"go-start/core/context"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"userId"`
	UserName  string `json:"userName"`
	UserPhone string `json:"userPhone"`
	Password  string `json:"password"`
	Base
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

func (u *User) GetByUserName(userName string) *User {
	if userName == "" {
		return nil
	}

	user := User{}
	err := context.Db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	return &user
}

func (u *User) Create(user *User) uint {
	var id uint
	err := context.Db.Transaction(func(tx *gorm.DB) error {
		create := UserIdCreate{}
		id = create.getNextUserId()

		user.CreatedByName = user.UserName
		user.CreatedBy = strconv.FormatUint(uint64(id), 10)
		user.UserId = id

		err := context.Db.Create(user).Error
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return id
}
