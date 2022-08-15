package models

import "go-start/core/context"

type UserIdCreate struct {
	Id  uint `json:"id"`
	Tub int
}

func (u UserIdCreate) getNextUserId() uint {
	create := UserIdCreate{}
	err := context.Db.Create(&create).Error
	if err != nil {
		panic(err)
	}

	return create.Id
}
