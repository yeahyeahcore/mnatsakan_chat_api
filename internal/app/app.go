package app

import (
	"log"
	"time"

	"mnatsakan_chat_api/internal/config"
	"mnatsakan_chat_api/internal/handler"
	"mnatsakan_chat_api/internal/repository"
	"mnatsakan_chat_api/internal/service"
	"mnatsakan_chat_api/pkg/database"
	"mnatsakan_chat_api/pkg/postgres"
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
		log.Fatalf("configuration file parsing error: %s", err.Error())
	}

	postgresDriver := postgres.GetDriver(&configuration.Database)

	databaseConnection, err := database.Connect(postgresDriver, loggerConfiguration)
	if err != nil {
		log.Fatalf("failed connected to DB: %s", err.Error())
	}

	repositoryInstance := repository.New(databaseConnection)
	serviceInstance := service.New(repositoryInstance)
	handlersInstance := handler.New(serviceInstance)
	serverInstance := new(server.Server)

	if err := serverInstance.Run(&configuration.HTTP, handlersInstance.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
