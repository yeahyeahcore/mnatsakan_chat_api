package repository

import (
	"fmt"

	"mnatsakan_chat_api/internal/model"

	"gorm.io/gorm"
)

// Message repository
type Message struct {
	database *gorm.DB
}

// Init repository instance
func (receiver *Message) Init(databaseConnection *gorm.DB) *Message {
	receiver.database = databaseConnection

	return receiver
}

// Insert element to messages table
func (receiver *Message) Insert(message *model.Message) (*model.Message, error) {
	if err := receiver.database.Table(messagesTable).Create(message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

// InsertMany elements to messages table
func (receiver *Message) InsertMany(messages *[]model.Message) (*[]model.Message, error) {
	return nil, fmt.Errorf("cannot add many messages in db")
}

// Find messages table elements containts arguments
func (receiver *Message) Find(arguments interface{}) (*[]model.Message, error) {
	var messages []model.Message

	if err := receiver.database.Table(messagesTable).Where(arguments).Find(&messages).Error; err != nil {
		return nil, err
	}

	return &messages, nil
}

// FindOne - find first messages table element containts arguments
func (receiver *Message) FindOne(arguments interface{}) (*model.Message, error) {
	var message model.Message

	if err := receiver.database.Table(messagesTable).Where(arguments).First(&message).Error; err != nil {
		return nil, err
	}

	return &message, nil
}

// FindByID - find messages table element by id
func (receiver *Message) FindByID(id string) (*model.Message, error) {
	var message model.Message

	if err := receiver.database.Table(messagesTable).First(&message, id).Error; err != nil {
		return nil, err
	}

	return &message, nil
}
