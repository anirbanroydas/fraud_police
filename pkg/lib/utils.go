// Package lib contains all helper functions required by every other package or file
// This packate does not depend on any other package hence no fear of circular dependency problems
package lib

import (
	"os"
)

// Getenv takes a key and a default string and tries to read the key from environment
// returns the actual value or else the fallback if not found in the environment
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
