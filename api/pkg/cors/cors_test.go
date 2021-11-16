package cors

import (
	"net/http"
	"testing"

	"github.com/rs/cors"
	"github.com/stretchr/testify/assert"
)

func TestCorsOptions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect cors.Options
	}{
		{
			name: "success",
			expect: cors.Options{
				AllowedOrigins: []string{
					"http://*",
					"https://*",
				},
				AllowedMethods: []string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
					http.MethodOptions,
				},
				AllowedHeaders: []string{
					"Accept",
					"Authorization",
					"Content-Type",
					"User-Agent",
					"X-Forwarded-For",
					"X-Forwarded-Proto",
					"X-Real-Ip",
				},
				AllowCredentials:   true,
				MaxAge:             1440, // 60m * 24h
				OptionsPassthrough: false,
				Debug:              false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewOptions())
		})
	}
}
