package service

import (
	"github.com/mnatsakan_chat_api/internal/dto"
	"github.com/mnatsakan_chat_api/internal/model"
)

// AuthService ...
type AuthService interface {
	Register(registerRequest dto.RegisterRequest) (*model.User, error)
	Login(loginRequest dto.LoginRequest) (*model.User, error)
}