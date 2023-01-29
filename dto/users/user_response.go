package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
