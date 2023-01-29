package usersdto

type CreateUserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
