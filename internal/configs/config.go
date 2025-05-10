package configs

import (
	"os"
)

// Config holds basic settings for the app, like the port and app name.
// You can load this from environment variables or use defaults.
type Config struct {
	AppName string
	Port    string
}

// LoadConfig reads environment variables and builds the config object.
// If PORT is not set, it falls back to the default (8080).
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not provided
	}

	return &Config{
		AppName: "Pack Calculator API", // can be used in logs or UI later
		Port:    port,
	}
}
