package repository

import (
	"gorm.io/gorm"
)

const (
	usersTable    = "users"
	messagesTable = "messages"
	chatsTable    = "chats"
)

// Repository is structure describes other repositories
type Repository struct {
	UserRepository
	MessageRepository
	ChatRepository
}

// New creates new instance of Repository
func New(databaseConnection *gorm.DB) *Repository {
	var userRepository User
	var messageRepository Message
	var chatRepository Chat

	return &Repository{
		UserRepository:    userRepository.Init(databaseConnection),
		MessageRepository: messageRepository.Init(databaseConnection),
		ChatRepository:    chatRepository.Init(databaseConnection),
	}
}
