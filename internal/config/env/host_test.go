package env_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestHost(t *testing.T) { //nolint:paralleltest
	type args func()

	tests := []test.Test[args, string]{
		{
			Name: "no env",
			Args: func() {},
			Want: "localhost",
		},
		{
			Name: "has env",
			Args: func() {
				t.Setenv("HOST", "google.com")
			},
			Want: "google.com",
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.Name, func(t *testing.T) {
			tt.Args()
			host := env.Host()

			require.Equal(t, tt.Want, host)
		})
	}
}
