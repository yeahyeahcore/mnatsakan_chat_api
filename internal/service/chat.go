package service

import (
	"mnatsakan_chat_api/internal/model"
	"mnatsakan_chat_api/internal/repository"
)

// Chat service
type Chat struct {
	repository repository.ChatRepository
}

// NewChatService - returns new chat service instance
func NewChatService(repository repository.ChatRepository) *Chat {
	return &Chat{
		repository: repository,
	}
}

// InsertChat - insert new chat
func (receiver *Chat) InsertChat(chat *model.Chat) (*model.Chat, error) {
	chat, err := receiver.repository.Insert(chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

// ReadAllUsersChats - reads all user's chats
func (receiver *Chat) ReadAllUsersChats(userID int64) (*[]model.Chat, error) {
	chats, err := receiver.repository.Find(map[string]interface{}{
		"user_id": userID,
	})

	return chats, err
}
