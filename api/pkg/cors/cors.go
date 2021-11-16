package cors

import (
	"net/http"

	"github.com/rs/cors"
)

func NewOptions() cors.Options {
	return cors.Options{
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
	}
}
