package api

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/validation"
	mock_validation "github.com/calmato/shs-web/api/mock/user/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var errmock = errors.New("some error")

type mocks struct {
	validator *mock_validation.MockRequestValidation
}

type testResponse struct {
	code codes.Code
	body proto.Message
}

type testOptions struct {
	now func() time.Time
}

type testOption func(opts *testOptions)
type grpcCaller func(ctx context.Context, service *userService) (proto.Message, error)

func newMocks(ctrl *gomock.Controller) *mocks {
	return &mocks{
		validator: mock_validation.NewMockRequestValidation(ctrl),
	}
}

func newUserService(mocks *mocks, opts *testOptions) *userService {
	return &userService{
		now:         opts.now,
		logger:      zap.NewNop(),
		sharedGroup: &singleflight.Group{},
		waitGroup:   &sync.WaitGroup{},
		validator:   mocks.validator,
	}
}

func testGRPC(
	setup func(context.Context, *testing.T, *mocks),
	expect *testResponse,
	grpcFn grpcCaller,
	opts ...testOption,
) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mocks := newMocks(ctrl)
		dopts := &testOptions{
			now: jst.Now,
		}
		for i := range opts {
			opts[i](dopts)
		}
		service := newUserService(mocks, dopts)
		ctx := context.Background()
		if setup != nil {
			setup(ctx, t, mocks)
		}
		res, err := grpcFn(ctx, service)
		if expect != nil {
			switch expect.code {
			case codes.OK:
				require.NoError(t, err)
			default:
				require.Error(t, err)
				status, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, expect.code, status.Code(), status.Code().String())
			}
			if expect.body != nil {
				require.Equal(t, expect.body, res)
			}
		}
		service.waitGroup.Wait()
	}
}

func TestUserService(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect user.UserServiceServer
	}{
		{name: "success"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.NotNil(t, NewUserService(&Params{}))
		})
	}
}

func TestGRPCError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		err    error
		expect codes.Code
	}{
		{
			name:   "error is nil",
			err:    nil,
			expect: codes.OK,
		},
		{
			name:   "invalid argument",
			err:    fmt.Errorf("%w: %s", validation.ErrRequestValidation, "test"),
			expect: codes.InvalidArgument,
		},
		{
			name:   "other error",
			err:    errmock,
			expect: codes.Unknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := gRPCError(tt.err)
			assert.Equal(t, tt.expect, status.Code(err))
		})
	}
}
