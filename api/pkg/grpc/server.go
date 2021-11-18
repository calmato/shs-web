package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer interface {
	Serve() error
	Stop()
}

type gRPCServer struct {
	server *grpc.Server
	lister net.Listener
}

func NewGRPCServer(server *grpc.Server, port int64) (GRPCServer, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("grpc: failed to listen port: %w", err)
	}
	return &gRPCServer{server: server, lister: lis}, nil
}

func (s *gRPCServer) Serve() error {
	return s.server.Serve(s.lister)
}

func (s *gRPCServer) Stop() {
	s.server.GracefulStop()
}
