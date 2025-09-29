package env

import (
	"os"
)

func Scheme() string {
	scheme := os.Getenv("SCHEME")

	if scheme == "" {
		scheme = "http"
	}

	return scheme
}
