package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name    string
		handler http.Handler
		port    int64
		expect  int
		isErr   bool
	}{
		{
			name: "success",
			handler: func() http.Handler {
				mux := http.NewServeMux()
				mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, "Hello World")
					w.WriteHeader(http.StatusOK)
				}))
				return mux
			}(),
			port:   20080,
			expect: http.StatusOK,
			isErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewHTTPServer(tt.handler, tt.port)
			go server.Serve()
			defer func() {
				server.Stop(ctx)
				time.Sleep(2 * time.Second)
			}()

			for i := 0; i < 10; i++ {
				time.Sleep(time.Microsecond * 100)
				url := fmt.Sprintf("http://localhost:%d/health", tt.port)
				res, err := http.Get(url)
				if err != nil && strings.Contains(err.Error(), "connect: connection refused") {
					continue
				}
				require.Equal(t, tt.isErr, err != nil, err)
				assert.Equal(t, tt.expect, res.StatusCode)
				break
			}
		})
	}
}
