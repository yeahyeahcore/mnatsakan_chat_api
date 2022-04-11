package service

import (
	"github.com/mnatsakan_chat_api/internal/dto"
	"github.com/mnatsakan_chat_api/internal/model"
	"github.com/mnatsakan_chat_api/internal/repository"
)

// Auth ...
type Auth struct {
	repository repository.UserRepository
}

// NewAuthService ...
func NewAuthService(repository repository.UserRepository) *Auth {
	return &Auth{repository: repository}
}

// Login ...
func (receiver *Auth) Login(loginRequest dto.LoginRequest) (*model.User, error) {
	// userRepository := receiver.repositories.UserRepository

	// loginRequest, err := utils.ParseJSON[domain.LoginRequest](request.Body)
	// if err != nil {
	// 	http.Error(response, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// findUserArguments := map[string]string{
	// 	"login": loginRequest.Login,
	// }

	// user, err := userRepository.FindOne(findUserArguments)
	// if err != nil {
	// 	http.Error(response, err.Error(), HTTPDatabaseErrorCodes[err])
	// 	return
	// }

	// tokenClaims, err := utils.GenerateToken(receiver.brancaInstance, strconv.FormatInt(user.ID, 10), user)
	// if err != nil {
	// 	http.Error(response, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// respond.Send(response, tokenClaims, http.StatusOK)
}

// Register ...
func (receiver *Auth) Register(registerRequest dto.RegisterRequest) (*model.User, error) {

}
