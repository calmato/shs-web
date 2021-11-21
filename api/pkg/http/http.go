package http

import (
	"context"
	"fmt"
	"net/http"
)

type Server interface {
	Serve() error
	Stop(ctx context.Context) error
}

type httpServer struct {
	server *http.Server
}

// NewHTTPServer - HTTPサーバーの生成
func NewHTTPServer(handler http.Handler, port int64) Server {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	return &httpServer{server: s}
}

func (s *httpServer) Serve() error {
	return s.server.ListenAndServe()
}

func (s *httpServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
