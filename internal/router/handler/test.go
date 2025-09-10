package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/yushafro/url-shortening-service/pkg/http"
	_test "github.com/yushafro/url-shortening-service/pkg/test"
)

type args http.ResponseArgs[string]
type want http.ResponseWant[string]
type test _test.Test[args, want]

type testServer struct {
	t      *testing.T
	server *httptest.Server
	test   test
}
