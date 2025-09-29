package env

import (
	"fmt"
	"os"
)

func URL() string {
	url := os.Getenv("URL")

	if url == "" {
		protocol := Scheme()
		host := Host()
		port := Port()

		return fmt.Sprintf("%s://%s:%s", protocol, host, port)
	}

	return url
}
