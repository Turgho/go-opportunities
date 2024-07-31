package main

import (
	"github.com/Turgho/go-opportunities.git/server"
	"github.com/Turgho/go-opportunities.git/settings"
)

var (
	logger *settings.Logger
)

func main() {
	logger = settings.GetLogger("main")

	// Initialize Logger
	err := settings.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Server
	server.Initialize()
}
