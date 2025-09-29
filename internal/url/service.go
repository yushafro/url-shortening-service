package url

import (
	"errors"

	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/httputils/url"
	"github.com/yushafro/url-shortening-service/pkg/id"
)

const idLength = 8

var (
	storage       = make(URLs) //nolint:gochecknoglobals
	ErrInvalidURL = errors.New("invalid URL")
)

func CutURL(str string) (string, error) {
	if !url.IsValidURL(str) {
		return "", ErrInvalidURL
	}

	id, _ := id.RandomID(idLength)
	storage[id] = str

	return env.URL() + "/" + id, nil
}
