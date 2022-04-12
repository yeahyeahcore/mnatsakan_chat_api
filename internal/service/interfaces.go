package service

import (
	"mnatsakan_chat_api/internal/dto"
)

// AuthService ...
type AuthService interface {
	Register(registerRequest *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(loginRequest *dto.LoginRequest) (*dto.LoginResponse, error)
}
