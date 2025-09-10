package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Port() string {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}
