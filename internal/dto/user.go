package dto

// User - structure of the returned entity from the request
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
