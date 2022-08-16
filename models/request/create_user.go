package request

type CreateUser struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required"`
	UserPhone string `json:"userPhone"`
}
