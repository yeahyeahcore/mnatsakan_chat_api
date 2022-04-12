package dto

// LoginRequest ...
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterRequest ...
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterResponse ...
type RegisterResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// LoginResponse ...
type LoginResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
