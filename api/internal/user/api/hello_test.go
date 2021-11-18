package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestHello(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.HelloRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().Hello(gomock.Any()).Return(nil)
			},
			req: &user.HelloRequest{
				Name: "test",
			},
			expect: &testResponse{
				code: codes.OK,
				body: &user.HelloResponse{
					Message: "Hello, test",
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().Hello(gomock.Any()).Return(validation.ErrRequestValidation)
			},
			req: &user.HelloRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *userService) (proto.Message, error) {
			return service.Hello(ctx, tt.req)
		}))
	}
}
