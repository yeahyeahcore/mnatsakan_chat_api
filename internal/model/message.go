package model

import "time"

// Message model
type Message struct {
	ID        int64     `json:"id" gorm:"column:id;primarykey;unique"`
	ChatID    int64     `json:"chatId" gorm:"column:chat_id"`
	UserID    int64     `json:"userId" gorm:"column:user_id"`
	Text      string    `json:"text" gorm:"column:text"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
	IsRead    bool      `json:"isRead" gorm:"column:is_read"`
}
