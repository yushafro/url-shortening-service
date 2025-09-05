package service

import (
	"fmt"

	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/model"
	"github.com/yushafro/url-shortening-service/internal/service/id"
)

var Urls = make(model.Urls)

func CutURL(url string) string {
	id := id.RandomID(8)

	Urls[id] = url

	return fmt.Sprintf("http://%s/%s", env.GetAddr(), id)
}
