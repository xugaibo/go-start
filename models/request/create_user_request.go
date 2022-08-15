package request

type CreateUserRequest struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required"`
	UserPhone string `json:"userPhone"`
}
