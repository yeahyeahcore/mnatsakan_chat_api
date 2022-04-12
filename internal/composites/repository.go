package composites

import (
	"fmt"
	"mnatsakan_chat_api/internal/config"
	"mnatsakan_chat_api/internal/repository"
	"mnatsakan_chat_api/pkg/database"
	"mnatsakan_chat_api/pkg/postgres"

	"gorm.io/gorm/logger"
)

// GetRepositoryInstance ...
func GetRepositoryInstance(config *config.DatabaseConfiguration, loggerConfig logger.Config) (*repository.Repository, error) {
	postgresDriver := postgres.GetDriver(config)

	databaseConnection, err := database.Connect(postgresDriver, loggerConfig)
	if err != nil {
		return nil, fmt.Errorf("failed connected to DB: %s", err.Error())
	}

	return repository.New(databaseConnection), nil
}
