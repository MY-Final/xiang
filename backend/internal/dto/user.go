package dto

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=32"`
	Password  string `json:"password" binding:"required,min=6,max=32"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Birthday  string `json:"birthday"`                   // YYYY-MM-DD
	Anniversary string `json:"anniversary"`             // YYYY-MM-DD
}

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserProfile 用户资料
type UserProfile struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
	Anniversary string `json:"anniversary"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	Token     string      `json:"token"`
	ExpiresIn int         `json:"expires_in"`
	User      UserProfile `json:"user"`
}
