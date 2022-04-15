package model

// Chat model
type Chat struct {
	ID     int64  `json:"id" gorm:"column:id;primarykey;unique"`
	UserID int64  `json:"userId" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
}
