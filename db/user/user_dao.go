package user

import (
	"errors"
	"go-start/core"
	"go-start/db"
	"gorm.io/gorm"
	"strconv"
)

type Dao struct {
	db.BaseDao
}

func NewDao(api core.Api) *Dao {
	dao := Dao{}
	api.MakeBaseDao(&dao.BaseDao)
	return &dao
}

func (u *Dao) GetByUserName(userName string) *User {
	if userName == "" {
		return nil
	}

	user := User{}
	err := u.Db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	return &user
}

func (u *Dao) Create(user *User) uint {
	var id uint
	err := u.Db.Transaction(func(tx *gorm.DB) error {
		id = getNextUserId(tx)

		user.CreatedByName = user.UserName
		user.CreatedBy = strconv.FormatUint(uint64(id), 10)
		user.UserId = id

		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return id
}

func getNextUserId(db *gorm.DB) uint {
	create := IdCreate{}
	err := db.Create(&create).Error
	if err != nil {
		panic(err)
	}

	return create.Id
}
