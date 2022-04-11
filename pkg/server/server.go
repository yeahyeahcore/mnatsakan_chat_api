package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mnatsakan_chat_api/pkg/types"
)

// Server ...
type Server struct {
	httpServer *http.Server
}

// Run ...
func (receiver *Server) Run(serverConfiguration types.HTTPConfiguration, handler http.Handler) error {
	serverAddress := fmt.Sprintf("%s:%s", serverConfiguration.Host, serverConfiguration.Port)

	receiver.httpServer = &http.Server{
		Addr:           serverAddress,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return receiver.httpServer.ListenAndServe()
}

// Shutdown ...
func (receiver *Server) Shutdown(ctx context.Context) error {
	return receiver.httpServer.Shutdown(ctx)
}
