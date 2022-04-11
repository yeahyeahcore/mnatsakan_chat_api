package database

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect to PostgreSQL database using GORM and return connection instance
func Connect(driver gorm.Dialector, loggerConfiguration logger.Config) (*gorm.DB, error) {
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), loggerConfiguration)

	databaseConnection, err := gorm.Open(driver, &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}

	return databaseConnection, nil
}
