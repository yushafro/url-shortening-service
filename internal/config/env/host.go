package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Host() string {
	godotenv.Load()

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	return host
}
