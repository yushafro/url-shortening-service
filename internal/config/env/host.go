package env

import (
	"os"
)

func Host() string {
	host := os.Getenv("HOST")

	if host == "" {
		host = "localhost"
	}

	return host
}
