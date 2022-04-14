package handler

import (
	"mnatsakan_chat_api/internal/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (receiver *Handler) getAuthMiddlewareConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		TokenLookup: "header:" + echo.HeaderAuthorization,
		ParseTokenFunc: func(accessToken string, _ echo.Context) (interface{}, error) {
			return receiver.services.AuthService.ParseToken(accessToken)
		},
		AuthScheme:    "Bearer",
		ContextKey:    userIDContextKey,
		SigningMethod: receiver.codec.SignMethod.Name,
		Claims:        dto.TokenClaims{},
	}
}
