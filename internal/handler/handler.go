package handler

import (
	"mnatsakan_chat_api/internal/service"
	"mnatsakan_chat_api/pkg/codec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	userIDContextKey = "user_id"
)

// Handler ...
type Handler struct {
	services *service.Service
	codec    *codec.Codec
}

// New ...
func New(services *service.Service, codec *codec.Codec) *Handler {
	return &Handler{
		services: services,
		codec:    codec,
	}
}

// InitRoutes ...
func (receiver Handler) InitRoutes() *echo.Echo {
	echo := echo.New()

	authGroup := echo.Group("/auth")
	{
		authGroup.POST("/login", receiver.login)
		authGroup.POST("/register", receiver.register)
	}

	privateGroup := echo.Group("/admin")
	{
		privateGroup.Use(middleware.JWTWithConfig(receiver.getAuthMiddlewareConfig()))

	}

	return echo
}
