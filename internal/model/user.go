package model

// User model
type User struct {
	ID       int64  `json:"id" gorm:"column:id;primarykey;unique"`
	Email    string `json:"email" gorm:"email:id;unique"`
	Username string `json:"username" gorm:"username:id"`
	Login    string `json:"login" gorm:"login:id;unique"`
	Password string `json:"password" gorm:"password:id"`
}
