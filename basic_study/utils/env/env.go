// Package env providers a set of environment variable related utilities.
package env

import "os"

// Get gets the value from multiple env variables.
func Get(names ...string) string {
	for _, n := range names {
		if v := os.Getenv(n); v != "" {
			return v
		}
	}

	return ""
}
