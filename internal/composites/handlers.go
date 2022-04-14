package composites

import (
	"mnatsakan_chat_api/internal/config"
	"mnatsakan_chat_api/internal/handler"
	"mnatsakan_chat_api/internal/service"

	"gorm.io/gorm/logger"
)

// GetHandlersInstance - get instance of handlers
func GetHandlersInstance(config *config.Config, loggerConfig logger.Config) (*handler.Handler, error) {
	repositoryInstance, err := GetRepositoryInstance(&config.Database, loggerConfig)
	if err != nil {
		return nil, err
	}

	codecInstance := GetCodecInstance()
	serviceInstance := service.New(repositoryInstance, codecInstance)

	return handler.New(serviceInstance, codecInstance), nil
}
