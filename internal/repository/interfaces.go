package repository

import (
	"mnatsakan_chat_api/internal/model"
	"mnatsakan_chat_api/pkg/repository"
)

// UserRepository - interface describes user's model responsibilities
type UserRepository interface {
	repository.IWrite[model.User]
	repository.IRead[model.User]
}

// MessageRepository - interface describes messages model responsibilities
type MessageRepository interface {
	repository.IWrite[model.Message]
	repository.IRead[model.Message]
}

// ChatRepository - interface describes chat's model responsibilities
type ChatRepository interface {
	repository.IWrite[model.Chat]
	repository.IRead[model.Chat]
}
