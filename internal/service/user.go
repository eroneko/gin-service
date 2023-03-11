package service

type CreateUserRequest struct {
	Username string `form:"username" binding:"required,min=5,max=100"`
	Password string `form:"password" binding:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
	Username  string `form:"username" binding:"omitempty,min=5,max=100"`
	Password  string `form:"password" binding:"omitempty,min=8,max=100"`
	Nickname  string `form:"nickname" binding:"omitempty,min=5,max=100"`
	AvatarURL string `form:"avatar" binding:"omitempty,min=5,max=100"`
}

type LoginRequest struct {
	Username string `form:"username" binding:"required,min=5,max=100"`
	Password string `form:"password" binding:"required,min=8,max=100"`
}

type GetUserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar"`
}
