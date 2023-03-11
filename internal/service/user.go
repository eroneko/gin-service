package service

type CreateUserRequest struct {
	UserName string `form:"username" binding:"required,min=5,max=100"`
	Password string `form:"password" binding:"required,min=8,max=100"`
}

type UpdateUserRequest struct {
	UserName  string `form:"username" binding:"omitempty,min=5,max=100"`
	Password  string `form:"password" binding:"omitempty,min=8,max=100"`
	NickName  string `form:"nickname" binding:"omitempty,min=5,max=100"`
	AvatarURL string `form:"avatar" binding:"omitempty,min=5,max=100"`
}

type LoginRequest struct {
	UserName string `form:"username" binding:"required,min=5,max=100"`
	Password string `form:"password" binding:"required,min=8,max=100"`
}

type GetUserResponse struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	AvatarURL string `json:"avatar"`
}
