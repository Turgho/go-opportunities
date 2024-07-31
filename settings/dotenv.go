package settings

import (
	"fmt"
	"os"
)

func GetEnv() ([]string, error) {
	// Find all variables
	env := []string{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	}

	// Verify all variables
	for _, value := range env {
		if value == "" {
			return nil, fmt.Errorf("missing .env variables")
		}
	}

	return env, nil
}
