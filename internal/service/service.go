package service

import (
	"mnatsakan_chat_api/internal/repository"
	"mnatsakan_chat_api/pkg/codec"
)

var (
	brancaSuperSecretKey = "supersecretbrancakey"
)

// Service contains the core logic.
type Service struct {
	AuthService
	MessageService
	ChatService
}

// New service instance
func New(repository *repository.Repository, codec *codec.Codec) *Service {
	return &Service{
		AuthService:    NewAuthService(repository.UserRepository, codec),
		MessageService: NewMessageService(repository.MessageRepository),
		ChatService:    NewChatService(repository.ChatRepository),
	}
}
