package repository

import (
	"fmt"

	"mnatsakan_chat_api/internal/model"

	"gorm.io/gorm"
)

// Chat repository
type Chat struct {
	database *gorm.DB
}

// Init repository instance
func (receiver *Chat) Init(databaseConnection *gorm.DB) *Chat {
	receiver.database = databaseConnection

	return receiver
}

// Insert element to chats table
func (receiver *Chat) Insert(chat *model.Chat) (*model.Chat, error) {
	if err := receiver.database.Table(chatsTable).Create(chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
}

// InsertMany elements to chats table
func (receiver *Chat) InsertMany(chat *[]model.Chat) (*[]model.Chat, error) {
	return nil, fmt.Errorf("cannot add many chats in db")
}

// Find chats table elements containts arguments
func (receiver *Chat) Find(arguments interface{}) (*[]model.Chat, error) {
	var chats []model.Chat

	if err := receiver.database.Table(chatsTable).Where(arguments).Find(&chats).Error; err != nil {
		return nil, err
	}

	return &chats, nil
}

// FindOne - find first chats table element containts arguments
func (receiver *Chat) FindOne(arguments interface{}) (*model.Chat, error) {
	var chat model.Chat

	if err := receiver.database.Table(messagesTable).Where(arguments).First(&chat).Error; err != nil {
		return nil, err
	}

	return &chat, nil
}

// FindByID - find chats table element by id
func (receiver *Chat) FindByID(id string) (*model.Chat, error) {
	var chat model.Chat

	if err := receiver.database.Table(messagesTable).First(&chat, id).Error; err != nil {
		return nil, err
	}

	return &chat, nil
}
