package url

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yushafro/url-shortening-service/pkg/httputils"
)

const (
	CutURLPattern = "/"
)

var (
	ErrRequiredURL           = errors.New("required URL")
	ErrContentTypeNotAllowed = errors.New("Content-Type is not allowed")
	ErrNoFoundByID           = errors.New("no found by ID")
)

func CutURLHandler(context *gin.Context) {
	if context.Request.Header.Get(httputils.ContentType) != httputils.URLEncoded {
		context.String(http.StatusBadRequest, ErrContentTypeNotAllowed.Error())

		return
	}

	urls, ok := context.GetPostFormArray("url")
	if !ok {
		context.String(http.StatusBadRequest, ErrRequiredURL.Error())

		return
	}

	urlIDs := make([]string, 0, len(urls))
	for _, u := range urls {
		id, err := CutURL(u)
		if err == nil {
			urlIDs = append(urlIDs, id)
		}
	}

	if len(urlIDs) == 0 {
		context.String(http.StatusBadRequest, ErrInvalidURL.Error())

		return
	}

	context.String(http.StatusCreated, strings.Join(urlIDs, "\n"))
}
