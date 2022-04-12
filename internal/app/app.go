package app

import (
	"log"
	"time"

	"mnatsakan_chat_api/internal/composites"
	"mnatsakan_chat_api/internal/config"
	"mnatsakan_chat_api/pkg/server"
	"mnatsakan_chat_api/pkg/utils"

	"gorm.io/gorm/logger"
)

var loggerConfiguration = logger.Config{
	SlowThreshold:             time.Second,   // Slow SQL threshold
	LogLevel:                  logger.Silent, // Log level
	IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	Colorful:                  false,         // Disable color
}

// Run app
func Run() {
	configuration, err := utils.ParseENV[config.Config]()
	if err != nil {
		log.Fatalf(err.Error())
	}

	handlersInstance, err := composites.GetHandlersInstance(configuration, loggerConfiguration)
	if err != nil {
		log.Fatalf(err.Error())
	}

	serverInstance := new(server.Server)

	if err := serverInstance.Run(&configuration.HTTP, handlersInstance.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
