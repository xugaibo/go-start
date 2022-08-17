package request

type UpdateToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
