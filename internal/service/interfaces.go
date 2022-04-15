package service

import (
	"mnatsakan_chat_api/internal/dto"
	"mnatsakan_chat_api/internal/model"
)

// AuthService - interface describes auth service responsibilities
type AuthService interface {
	Register(registerRequest *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
	ParseToken(token string) (string, error)
}

// MessageService - interface describes message service responsibilities
type MessageService interface {
	InsertMessage(messageModel *model.Message) (*model.Message, error)
	ReadAllChatMessages(chatID int64) (*[]model.Message, error)
}

// ChatService - interface describes chat service responsibilities
type ChatService interface {
	InsertChat(chatModel *model.Chat) (*model.Chat, error)
	ReadAllUsersChats(userID int64) (*[]model.Chat, error)
}
