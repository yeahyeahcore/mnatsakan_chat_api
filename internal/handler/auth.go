package handler

import (
	"net/http"

	"github.com/mnatsakan_chat_api/internal/dto"
	"github.com/mnatsakan_chat_api/pkg/utils"
)

func (receiver *Handler) register(request *http.Request, response http.ResponseWriter) {
	defer request.Body.Close()

	// registerInputData, err := utils.ParseJSON[handler_types.LoginRequest](request.Body)
	// if err != nil {
	// 	http.Error(response, err.Error(), http.StatusBadRequest)
	// 	return
	// }
}

func (receiver *Handler) login(request *http.Request, response http.ResponseWriter) {
	defer request.Body.Close()

	loginInput, err := utils.ParseJSON[dto.LoginRequest](request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

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
