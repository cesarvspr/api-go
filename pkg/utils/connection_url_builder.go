package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	url := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	// Return connection URL.
	return url, nil
}
