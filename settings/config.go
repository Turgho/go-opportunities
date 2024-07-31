package settings

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize DB
	db, err = IniatializePostgres()

	if err != nil {
		return fmt.Errorf("error initialize postgres database: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
