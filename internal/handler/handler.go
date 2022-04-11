package handler

import (
	"github.com/mnatsakan_chat_api/internal/service"

	"github.com/go-martini/martini"
)

// Handler ...
type Handler struct {
	services *service.Service
}

// New ...
func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// InitRoutes ...
func (receiver Handler) InitRoutes() *martini.ClassicMartini {
	martiniInstance := martini.Classic()

	martiniInstance.Get("/login", receiver.login)
	martiniInstance.Post("/register", receiver.register)

	return martiniInstance
}
