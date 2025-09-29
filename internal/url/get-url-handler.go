package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushafro/url-shortening-service/pkg/httputils"
)

const GetURLPattern = "/:id"

func GetURLHandler(context *gin.Context) {
	id := context.Param("id")

	originURL := storage[id]

	if originURL == "" {
		context.String(http.StatusBadRequest, ErrNoFoundByID.Error())

		return
	}

	context.Header(httputils.Location, originURL)
	context.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
