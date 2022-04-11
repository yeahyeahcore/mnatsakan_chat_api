package service

import (
	"github.com/mnatsakan_chat_api/internal/repository"
)

var (
	brancaSuperSecretKey = "supersecretbrancakey"
)

// Service contains the core logic.
type Service struct {
	AuthService
}

// New service instance
func New(repository *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repository),
	}
}
