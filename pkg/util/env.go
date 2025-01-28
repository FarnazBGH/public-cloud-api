package util

import (
	"fmt"
	"os"
)

// LoadEnv loads and validates required environment variables
func LoadEnv() (string, string) {
	apiKey := os.Getenv("API_KEY")
	apiHost := os.Getenv("API_HOST")
	if apiKey == "" || apiHost == "" {
		fmt.Fprintln(os.Stderr, "Error: API_KEY and API_HOST environment variables must be set")
		os.Exit(1)
	}
	return apiKey, apiHost
}
