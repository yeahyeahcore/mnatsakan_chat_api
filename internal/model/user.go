package model

// User model
type User struct {
	ID       int64  `json:"id" gorm:"column:id;primarykey;unique"`
	Email    string `json:"email" gorm:"column:email;unique"`
	Username string `json:"username" gorm:"column:username"`
	Login    string `json:"login" gorm:"column:login;unique"`
	Password string `json:"password" gorm:"column:password"`
}
