package env

import (
	"fmt"
)

func URL() string {
	protocol := Protocol()
	host := Host()
	port := Port()

	return fmt.Sprintf("%s://%s:%s", protocol, host, port)
}
