package handler

import (
	"net/http"
	"strconv"

	"mnatsakan_chat_api/internal/handler/respond"
	"mnatsakan_chat_api/internal/model"
	"mnatsakan_chat_api/pkg/utils"

	"github.com/labstack/echo/v4"
)

const (
	chatIDParam = "chat_id"
)

func (receiver *Handler) readAllChatMessages(context echo.Context) error {
	chatID := context.QueryParam(chatIDParam)

	isNumber := utils.IsNumber(chatID)
	if !isNumber || chatID == "" {
		return respond.SendMessage(context, "chat_id is not a number", http.StatusBadRequest)
	}

	parsedChatID, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	messages, err := receiver.services.MessageService.ReadAllChatMessages(parsedChatID)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	return context.JSON(http.StatusOK, messages)
}

func (receiver *Handler) insertMessage(context echo.Context) error {
	defer context.Request().Body.Close()

	message, err := utils.ParseJSON[model.Message](context.Request().Body)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	createdMessage, err := receiver.services.MessageService.InsertMessage(message)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	return context.JSON(http.StatusOK, createdMessage)
}
