package api

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	mock_database "github.com/calmato/shs-web/api/mock/lesson/database"
	mock_validation "github.com/calmato/shs-web/api/mock/lesson/validation"
	mock_classroom "github.com/calmato/shs-web/api/mock/proto/classroom"
	mock_messenger "github.com/calmato/shs-web/api/mock/proto/messenger"
	mock_user "github.com/calmato/shs-web/api/mock/proto/user"
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
	classroom *mock_classroom.MockClassroomServiceClient
	user      *mock_user.MockUserServiceClient
	messenger *mock_messenger.MockMessengerServiceClient
}

type dbMocks struct {
	Lesson               *mock_database.MockLesson
	ShiftSummary         *mock_database.MockShiftSummary
	Shift                *mock_database.MockShift
	TeacherSubmission    *mock_database.MockTeacherSubmission
	TeacherShift         *mock_database.MockTeacherShift
	StudentSubmission    *mock_database.MockStudentSubmission
	StudentShift         *mock_database.MockStudentShift
	StudentShiftTemplate *mock_database.MockStudentShiftTemplate
}

type testResponse struct {
	code codes.Code
	body proto.Message
}

type testOptions struct {
	now func() time.Time
}

type testOption func(opts *testOptions)
type grpcCaller func(ctx context.Context, service *lessonService) (proto.Message, error)

func newMocks(ctrl *gomock.Controller) *mocks {
	return &mocks{
		db:        newDBMocks(ctrl),
		validator: mock_validation.NewMockRequestValidation(ctrl),
		classroom: mock_classroom.NewMockClassroomServiceClient(ctrl),
		user:      mock_user.NewMockUserServiceClient(ctrl),
		messenger: mock_messenger.NewMockMessengerServiceClient(ctrl),
	}
}

func newDBMocks(ctrl *gomock.Controller) *dbMocks {
	return &dbMocks{
		Lesson:               mock_database.NewMockLesson(ctrl),
		ShiftSummary:         mock_database.NewMockShiftSummary(ctrl),
		Shift:                mock_database.NewMockShift(ctrl),
		TeacherSubmission:    mock_database.NewMockTeacherSubmission(ctrl),
		TeacherShift:         mock_database.NewMockTeacherShift(ctrl),
		StudentSubmission:    mock_database.NewMockStudentSubmission(ctrl),
		StudentShift:         mock_database.NewMockStudentShift(ctrl),
		StudentShiftTemplate: mock_database.NewMockStudentShiftTemplate(ctrl),
	}
}

func newLessonService(mocks *mocks, opts *testOptions) *lessonService {
	return &lessonService{
		now:         opts.now,
		logger:      zap.NewNop(),
		sharedGroup: &singleflight.Group{},
		waitGroup:   &sync.WaitGroup{},
		db: &database.Database{
			Lesson:               mocks.db.Lesson,
			ShiftSummary:         mocks.db.ShiftSummary,
			Shift:                mocks.db.Shift,
			TeacherSubmission:    mocks.db.TeacherSubmission,
			TeacherShift:         mocks.db.TeacherShift,
			StudentSubmission:    mocks.db.StudentSubmission,
			StudentShift:         mocks.db.StudentShift,
			StudentShiftTemplate: mocks.db.StudentShiftTemplate,
		},
		validator: mocks.validator,
		classroom: mocks.classroom,
		user:      mocks.user,
		messenger: mocks.messenger,
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
		service := newLessonService(mocks, dopts)
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

func TestLessonService(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, NewLessonService(&Params{}))
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
			name:   "grpc error",
			err:    status.Error(codes.Unavailable, "test"),
			expect: codes.Unavailable,
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
