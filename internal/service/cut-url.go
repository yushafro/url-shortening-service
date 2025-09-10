package service

import (
	"fmt"

	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/model"
	"github.com/yushafro/url-shortening-service/pkg/http"
	"github.com/yushafro/url-shortening-service/pkg/http/url"
	"github.com/yushafro/url-shortening-service/pkg/id"
)

var URLs = make(model.Urls)
var ErrInvalidURL = fmt.Errorf(http.Invalid, "URL")

func CutURL(s string) (string, error) {
	if !url.IsValidURL(s) {
		return "", ErrInvalidURL
	}

	id, _ := id.RandomID(8)
	URLs[id] = s

	return env.URL() + "/" + id, nil
}
