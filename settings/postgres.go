package settings

import (
	"github.com/Turgho/go-opportunities.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IniatializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Database connection
	psql := "host=api_db user=postgres password=password dbname=opportunities_db port=5432 sslmode=disable"

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
