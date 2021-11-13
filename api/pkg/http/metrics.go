package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsServer struct {
	server *http.Server
}

// NewMetricsServer - メトリクス取得用サーバーの生成
func NewMetricsServer(port int64) HTTPServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return &metricsServer{server: s}
}

func (s *metricsServer) Serve() error {
	return s.server.ListenAndServe()
}

func (s *metricsServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
