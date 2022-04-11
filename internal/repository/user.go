package repository

import (
	"github.com/mnatsakan_chat_api/internal/model"

	"gorm.io/gorm"
)

// User repository
type User struct {
	database *gorm.DB
}

// Init repository instance
func (receiver *User) Init(databaseConnection *gorm.DB) *User {
	receiver.database = databaseConnection

	return receiver
}

// Insert element to users table
func (receiver *User) Insert(arguments interface{}) (*model.User, error) {
	return &model.User{}, nil
}

// InsertMany elements to users table
func (receiver *User) InsertMany(arguments []interface{}) (*[]model.User, error) {
	return &[]model.User{}, nil
}

// Find users table elements containts arguments
func (receiver *User) Find(arguments interface{}) (*[]model.User, error) {
	return &[]model.User{}, nil
}

// FindOne - find first users table element containts arguments
func (receiver *User) FindOne(arguments interface{}) (*model.User, error) {
	var user model.User

	if err := receiver.database.Table(usersTable).Where(arguments).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByID - find users table element by id
func (receiver *User) FindByID(id string) (*model.User, error) {
	var user model.User

	if err := receiver.database.Table(usersTable).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
