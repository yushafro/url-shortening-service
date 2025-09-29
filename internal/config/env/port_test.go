package env_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestPort(t *testing.T) { //nolint:paralleltest
	type args func()

	tests := []test.Test[args, string]{
		{
			Name: "no env",
			Args: func() {},
			Want: "3000",
		},
		{
			Name: "has env",
			Args: func() {
				t.Setenv("PORT", "7000")
			},
			Want: "7000",
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.Name, func(t *testing.T) {
			tt.Args()
			port := env.Port()

			require.Equal(t, tt.Want, port)
		})
	}
}
