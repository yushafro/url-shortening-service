package env

import (
	"os"
)

func Port() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return port
}
