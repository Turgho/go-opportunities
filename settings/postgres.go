package settings

import (
	"fmt"

	"github.com/Turgho/go-opportunities.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IniatializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Get .env variables
	env, err := GetEnv()
	if err != nil {
		logger.Errorf("can't get .env variables: %v", err)
		return nil, err
	}

	// Database connection
	psql := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		env[0], env[1], env[2], env[3],
	)

	// Conect to postgres database
	db, err := gorm.Open(postgres.Open(psql), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to the database: %v", err)
		return nil, err
	}

	// Migration schemas
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("postgres failed to migrate: %v", err)
	}

	// Return DB
	return db, nil
}
