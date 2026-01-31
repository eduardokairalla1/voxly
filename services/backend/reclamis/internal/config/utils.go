// * Utility functions for configuration management. *

package config

// --- IMPORTS ---
import (
	"os"
)

// --- CODE ---
// Get environment variable or return default value.
func getEnv(key, defaultValue string) string {

	// Get environment variable
	value := os.Getenv(key)

	// Value exists: return it
	if value != "" {
		return value
	}

	// Value does not exist: return default
	return defaultValue
}
