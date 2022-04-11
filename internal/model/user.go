package model

// User model
type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Login    string `json:"login,omitempty"`
}
