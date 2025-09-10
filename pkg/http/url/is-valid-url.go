package url

import "net/url"

func IsValidUrl(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	} else if u.Host == "" {
		return false
	}

	return true
}
