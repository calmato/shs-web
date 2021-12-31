package api

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/database"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	mock_database "github.com/calmato/shs-web/api/mock/classroom/database"
	mock_validation "github.com/calmato/shs-web/api/mock/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
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
	db        *dbMocks
	validator *mock_validation.MockRequestValidation
}

type dbMocks struct {
	Subject        *mock_database.MockSubject
	TeacherSubject *mock_database.MockTeacherSubject
	Schedule       *mock_database.MockSchedule
	Room           *mock_database.MockRoom
}

type testResponse struct {
	code codes.Code
	body proto.Message
}

type testOptions struct {
	now func() time.Time
}

type testOption func(opts *testOptions)
type grpcCaller func(ctx context.Context, service *classroomService) (proto.Message, error)

func newMocks(ctrl *gomock.Controller) *mocks {
	return &mocks{
		db:        newDBMocks(ctrl),
		validator: mock_validation.NewMockRequestValidation(ctrl),
	}
}

func newDBMocks(ctrl *gomock.Controller) *dbMocks {
	return &dbMocks{
		Subject:        mock_database.NewMockSubject(ctrl),
		TeacherSubject: mock_database.NewMockTeacherSubject(ctrl),
		Schedule:       mock_database.NewMockSchedule(ctrl),
		Room:           mock_database.NewMockRoom(ctrl),
	}
}

func newClassroomService(mocks *mocks, opts *testOptions) *classroomService {
	return &classroomService{
		now:         opts.now,
		logger:      zap.NewNop(),
		sharedGroup: &singleflight.Group{},
		waitGroup:   &sync.WaitGroup{},
		db: &database.Database{
			Subject:        mocks.db.Subject,
			TeacherSubject: mocks.db.TeacherSubject,
			Schedule:       mocks.db.Schedule,
			Room:           mocks.db.Room,
		},
		validator: mocks.validator,
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
		service := newClassroomService(mocks, dopts)
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

func TestClassroomService(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, NewClassroomService(&Params{}))
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
			name:   "not found",
			err:    fmt.Errorf("%w: %s", database.ErrNotFound, "test"),
			expect: codes.NotFound,
		},
		{
			name:   "already exists",
			err:    fmt.Errorf("%w: %s", database.ErrAlreadyExists, "test"),
			expect: codes.AlreadyExists,
		},
		{
			name:   "unimplemented",
			err:    fmt.Errorf("%w: %s", database.ErrNotImplemented, "test"),
			expect: codes.Unimplemented,
		},
		{
			name:   "other error",
			err:    errmock,
			expect: codes.Internal,
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
