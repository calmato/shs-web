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

func TestMetricsServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name   string
		port   int64
		expect int
		isErr  bool
	}{
		{
			name:   "success",
			port:   20081,
			expect: http.StatusOK,
			isErr:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			server := NewMetricsServer(tt.port)
			go server.Serve()
			defer func() {
				server.Stop(ctx)
				time.Sleep(2 * time.Second)
			}()

			for i := 0; i < 10; i++ {
				time.Sleep(time.Microsecond * 100)
				url := fmt.Sprintf("http://localhost:%d/metrics", tt.port)
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
