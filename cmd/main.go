package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdelrhman-elsbagh/pack-calculator/internal/configs"
	"github.com/abdelrhman-elsbagh/pack-calculator/internal/transport/http"
)

func main() {
	// Load app configuration (like port, app name.)
	config := configs.LoadConfig()

	// Set up the HTTP router with routes and middleware
	router := http.SetupRouter()

	// Print a friendly message to indicate the server is running
	fmt.Printf("App %s is running on port %s...\n", config.AppName, config.Port)

	// Start the server on the configured port
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
