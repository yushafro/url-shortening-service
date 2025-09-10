package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Protocol() string {
	godotenv.Load()

	protocol := os.Getenv("PROTOCOL")
	if protocol == "" {
		protocol = "http"
	}

	return protocol
}
