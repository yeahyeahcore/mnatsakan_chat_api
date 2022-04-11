package repository

import (
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

// Repository is structure describes other repositories
type Repository struct {
	UserRepository
}

// New creates new instance of Repository
func New(databaseConnection *gorm.DB) *Repository {
	var userRepository User

	return &Repository{
		UserRepository: userRepository.Init(databaseConnection),
	}
}
