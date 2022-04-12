package handler

import (
	"net/http"

	"mnatsakan_chat_api/internal/dto"
	"mnatsakan_chat_api/internal/handler/respond"
	"mnatsakan_chat_api/pkg/utils"
)

func (receiver *Handler) register(request *http.Request, response http.ResponseWriter) {
	defer request.Body.Close()

	registerInputData, err := utils.ParseJSON[dto.RegisterRequest](request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	registerResponse, err := receiver.services.AuthService.Register(registerInputData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	respond.SendJSON(response, registerResponse, http.StatusCreated)
}

func (receiver *Handler) login(request *http.Request, response http.ResponseWriter) {
	defer request.Body.Close()

	loginInput, err := utils.ParseJSON[dto.LoginRequest](request.Body)
	if err != nil {
		respond.SendMessage(response, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse, err := receiver.services.AuthService.Login(loginInput)
	if err != nil {
		respond.SendMessage(response, err.Error(), http.StatusInternalServerError)
		return
	}

	respond.SendJSON(response, loginResponse, http.StatusOK)
}
