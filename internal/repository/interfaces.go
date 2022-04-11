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
