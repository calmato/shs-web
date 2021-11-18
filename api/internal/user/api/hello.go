package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

func (s *userService) Hello(ctx context.Context, req *user.HelloRequest) (*user.HelloResponse, error) {
	if err := s.validator.Hello(req); err != nil {
		return nil, gRPCError(err)
	}
	res := &user.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}
	return res, nil
}
