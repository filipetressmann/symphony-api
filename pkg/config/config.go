package config

import "os"

// GetEnv retrieves the value of the environment variable named by key.
// If the variable is not set, it returns the defaultValue.
// This function is useful for providing default values for configuration settings
// in applications, allowing them to run with sensible defaults without requiring
// explicit configuration in every environment.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
