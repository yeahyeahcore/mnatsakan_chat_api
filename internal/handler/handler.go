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

	// Settings
	receiver.setRoutesSettings(echo)

	// Routes
	authGroup := echo.Group("/auth")
	{
		authGroup.POST("/login", receiver.login)
		authGroup.POST("/register", receiver.register)
	}

	privateGroup := echo.Group("")
	{
		privateGroup.Use(middleware.JWTWithConfig(receiver.getAuthMiddlewareConfig()))

		privateGroup.GET("/messages", receiver.readAllChatMessages)
		privateGroup.GET("/chats", receiver.readAllUsersChats)
		privateGroup.POST("/message", receiver.insertMessage)
		privateGroup.POST("/chat", receiver.insertChat)
	}

	return echo
}

func (receiver Handler) setRoutesSettings(echo *echo.Echo) {
	echo.Static("/media", "/static")
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
	}))
}
