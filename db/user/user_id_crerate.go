package user

type IdCreate struct {
	Id  uint `json:"id"`
	Tub int
}

func (IdCreate) TableName() string {
	return "user_id_create"
}
