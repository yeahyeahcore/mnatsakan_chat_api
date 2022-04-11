package postgres

import (
	"fmt"

	"mnatsakan_chat_api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDriver ...
func GetDriver(databaseConfiguration *config.DatabaseConfiguration) gorm.Dialector {
	connectionString := generateConnectionString(databaseConfiguration)

	return postgres.Open(connectionString)
}

func generateConnectionString(config *config.DatabaseConfiguration) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?ssl_mode=%s",
		config.User, config.Password, config.Host, config.Port, config.Database, config.SSLMode,
	)
}
