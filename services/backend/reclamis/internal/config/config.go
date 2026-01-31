// * Main functions for configuration management. *

package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// --- CODE ---
// Loads the configuration from environment variables.
// Returns a Config struct with the loaded values.
func Load() *Configs {

	// Load .env file if it exists
	_ = godotenv.Load()

	// Initialize default version
	version := "0.0.1"

	// Read version from VERSION file
	data, err := os.ReadFile("VERSION")

	// Read version successfully: set version from file
	if err == nil {
		version = strings.TrimSpace(string(data))
	}

	// Return configuration struct
	return &Configs{
		Port:     getEnv("RECLAMIS_PORT", "8080"),
		LogLevel: getEnv("RECLAMIS_LOG_LEVEL", "info"),
		Version:  version,
	}
}
