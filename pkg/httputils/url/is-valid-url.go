package url

import (
	"log/slog"
	"net/url"
)

func IsValidURL(s string) bool {
	parsedURL, err := url.Parse(s)
	if err != nil {
		slog.Error("pkg/httputils/url/is-valid-url: " + err.Error())

		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		slog.Error("pkg/httputils/url/is-valid-url: bad URL scheme - " + parsedURL.Scheme)

		return false
	} else if parsedURL.Host == "" {
		slog.Error("pkg/httputils/url/is-valid-url: bad URL host - " + parsedURL.Host)

		return false
	}

	return true
}
