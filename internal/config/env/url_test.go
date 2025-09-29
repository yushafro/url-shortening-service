package env_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestURL(t *testing.T) { //nolint:paralleltest
	type args func()
	tests := []test.Test[args, string]{
		{
			Name: "no env",
			Args: func() {
				t.Setenv("SCHEME", "https")
				t.Setenv("HOST", "yandex.ru")
				t.Setenv("PORT", "3000")
			},
			Want: "https://yandex.ru:3000",
		},
		{
			Name: "has env",
			Args: func() {
				t.Setenv("URL", "https://google.com")
			},
			Want: "https://google.com",
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.Name, func(t *testing.T) {
			tt.Args()
			url := env.URL()

			require.Equal(t, tt.Want, url)
		})
	}
}
