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
}

// New service instance
func New(repository *repository.Repository, codec *codec.Codec) *Service {
	return &Service{
		AuthService: NewAuthService(repository.UserRepository, codec),
	}
}
