package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"mnatsakan_chat_api/internal/config"
)

// Server ...
type Server struct {
	httpServer *http.Server
}

// Run ...
func (receiver *Server) Run(serverConfiguration *config.HTTPConfiguration, handler http.Handler) error {
	serverAddress := fmt.Sprintf("%s:%s", serverConfiguration.Host, serverConfiguration.Port)

	receiver.httpServer = &http.Server{
		Addr:           serverAddress,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	fmt.Printf("Listening and serving HTTP on %s\n", serverAddress)

	return receiver.httpServer.ListenAndServe()
}

// Shutdown ...
func (receiver *Server) Shutdown(ctx context.Context) error {
	return receiver.httpServer.Shutdown(ctx)
}
