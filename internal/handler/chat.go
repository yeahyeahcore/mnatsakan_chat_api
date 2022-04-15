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
	userIDParam = "user_id"
)

func (receiver *Handler) readAllUsersChats(context echo.Context) error {
	userID := context.QueryParam(userIDParam)

	isNumber := utils.IsNumber(userID)
	if !isNumber || userID == "" {
		return respond.SendMessage(context, "user_id is not a number", http.StatusBadRequest)
	}

	parsedUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	chats, err := receiver.services.ChatService.ReadAllUsersChats(parsedUserID)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	return context.JSON(http.StatusOK, chats)
}

func (receiver *Handler) insertChat(context echo.Context) error {
	defer context.Request().Body.Close()

	chat, err := utils.ParseJSON[model.Chat](context.Request().Body)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusBadRequest)
	}

	createdChat, err := receiver.services.ChatService.InsertChat(chat)
	if err != nil {
		return respond.SendMessage(context, err.Error(), http.StatusInternalServerError)
	}

	return context.JSON(http.StatusOK, createdChat)
}
