package url_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/url"
	"github.com/yushafro/url-shortening-service/pkg/httputils"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

type (
	request   test.Request[string]
	response  test.Response[string]
	tableTest test.Test[request, response]
)

type TestServer struct {
	T      *testing.T
	Server *httptest.Server
	Test   tableTest
}

func TestCutUrlHandler(t *testing.T) {
	t.Parallel()

	tests := []tableTest{
		{
			Name: "valid request",
			Args: request{
				Method:      http.MethodPost,
				Path:        "/",
				Body:        "url=https%3A%2F%2Fgoogle.com&url=https%3A%2F%2Fyandex.ru&url=bad%3Abad.ru",
				ContentType: httputils.URLEncoded,
			},
			Want: response{
				StatusCode: http.StatusCreated,
				Response:   fmt.Sprintf("%s/11223344\n%s/11223344", env.URL(), env.URL()),
			},
		},
		{
			Name: "empty body",
			Args: request{
				Method:      http.MethodPost,
				Path:        "/",
				Body:        "",
				ContentType: httputils.URLEncoded,
			},
			Want: response{
				StatusCode: http.StatusBadRequest,
				Response:   url.ErrRequiredURL.Error(),
				WantError:  true,
			},
		},
		{
			Name: "invalid Content-Type",
			Args: request{
				Method:      http.MethodPost,
				Path:        "/",
				Body:        "url=https://google.com",
				ContentType: httputils.JSON,
			},
			Want: response{
				StatusCode: http.StatusBadRequest,
				Response:   url.ErrContentTypeNotAllowed.Error(),
				WantError:  true,
			},
		},
		{
			Name: "invalid URL",
			Args: request{
				Method:      http.MethodPost,
				Path:        "/",
				Body:        "url=htp:/google.com",
				ContentType: httputils.URLEncoded,
			},
			Want: response{
				StatusCode: http.StatusBadRequest,
				Response:   url.ErrInvalidURL.Error(),
				WantError:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.POST(url.CutURLPattern, url.CutURLHandler)
			server := httptest.NewServer(router)
			defer server.Close()

			testServer := &TestServer{
				T:      t,
				Server: server,
				Test:   test,
			}

			body := postID(testServer)

			if testServer.Test.Want.WantError {
				require.Equal(testServer.T, testServer.Test.Want.Response, body)

				return
			}

			require.Len(testServer.T, body, len(testServer.Test.Want.Response))
		})
	}
}

func postID(testServer *TestServer) string {
	res := req.DevMode().
		Post(testServer.Server.URL).
		SetHeader(httputils.ContentType, testServer.Test.Args.ContentType).
		SetBodyString(testServer.Test.Args.Body).
		Do()

	require.NoError(testServer.T, res.Err)
	body := strings.TrimSpace(res.String())

	return body
}
