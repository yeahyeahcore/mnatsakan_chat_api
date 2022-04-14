package handler

import (
	"fmt"
	"net/http"

	"mnatsakan_chat_api/internal/dto"
	"mnatsakan_chat_api/internal/handler/respond"
	"mnatsakan_chat_api/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (receiver *Handler) register(context echo.Context) error {
	defer context.Request().Body.Close()

	registerInputData, err := utils.ParseJSON[dto.RegisterRequest](context.Request().Body)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	registerResponse, err := receiver.services.AuthService.Register(registerInputData)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	return context.JSON(http.StatusCreated, registerResponse)
}

func (receiver *Handler) login(context echo.Context) error {
	defer context.Request().Body.Close()

	loginInput, err := utils.ParseJSON[dto.LoginRequest](context.Request().Body)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	loginResponse, err := receiver.services.AuthService.Login(loginInput)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	context.Response().Header().Set("Set-Cookie", fmt.Sprintf("token=%s;HttpOnly", loginResponse.Token))

	return context.JSON(http.StatusOK, loginResponse)
}
