package service

import (
	"mnatsakan_chat_api/internal/model"
	"mnatsakan_chat_api/internal/repository"
)

// Message service
type Message struct {
	repository repository.MessageRepository
}

// NewMessageService - returns new message service instance
func NewMessageService(repository repository.MessageRepository) *Message {
	return &Message{
		repository: repository,
	}
}

// InsertMessage - insert new message
func (receiver *Message) InsertMessage(message *model.Message) (*model.Message, error) {
	message, err := receiver.repository.Insert(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// ReadAllChatMessages - reads all messages of chat
func (receiver *Message) ReadAllChatMessages(chatID int64) (*[]model.Message, error) {
	messages, err := receiver.repository.Find(map[string]interface{}{
		"chat_id": chatID,
	})

	return messages, err
}
