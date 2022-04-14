package respond

import (
	"github.com/labstack/echo/v4"
)

// SendMessage sends message response to client
func SendMessage(context echo.Context, message string, statusCode int) error {
	messageData := map[string]string{
		"message": message,
	}

	return context.JSON(statusCode, messageData)
}
