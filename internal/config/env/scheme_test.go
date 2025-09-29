package env_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestScheme(t *testing.T) { //nolint:paralleltest
	type args func()

	tests := []test.Test[args, string]{
		{
			Name: "no env",
			Args: func() {},
			Want: "http",
		},
		{
			Name: "has env",
			Args: func() {
				t.Setenv("SCHEME", "https")
			},
			Want: "https",
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.Name, func(t *testing.T) {
			tt.Args()
			scheme := env.Scheme()

			require.Equal(t, tt.Want, scheme)
		})
	}
}
