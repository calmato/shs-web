package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListSchedules(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.ListSchedulesRequest{}
	schedules := entity.Schedules{
		{
			Weekday:   time.Sunday,
			IsClosed:  true,
			Lessons:   nil,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Weekday:  time.Monday,
			IsClosed: false,
			Lessons: entity.Lessons{
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.ListSchedulesRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListSchedules(req).Return(nil)
				mocks.db.Schedule.EXPECT().List(gomock.Any()).Return(schedules, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.ListSchedulesResponse{
					Schedules: []*classroom.Schedule{
						{
							Weekday:   int32(time.Sunday),
							IsClosed:  true,
							Lessons:   []*classroom.Schedule_Lesson{},
							CreatedAt: now.Unix(),
							UpdatedAt: now.Unix(),
						},
						{
							Weekday:  int32(time.Monday),
							IsClosed: true,
							Lessons: []*classroom.Schedule_Lesson{
								{StartTime: "1700", EndTime: "1830"},
								{StartTime: "1830", EndTime: "2000"},
							},
							CreatedAt: now.Unix(),
							UpdatedAt: now.Unix(),
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.ListSchedulesRequest{}
				mocks.validator.EXPECT().ListSchedules(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.ListSchedulesRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list schedules",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListSchedules(req).Return(nil)
				mocks.db.Schedule.EXPECT().List(gomock.Any()).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *classroomService) (proto.Message, error) {
			return service.ListSchedules(ctx, tt.req)
		}))
	}
}

func TestGetSchedule(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.GetScheduleRequest{
		Weekday: int32(time.Monday),
	}
	schedule := &entity.Schedule{
		Weekday:  time.Monday,
		IsClosed: false,
		Lessons: entity.Lessons{
			{StartTime: "1700", EndTime: "1830"},
			{StartTime: "1830", EndTime: "2000"},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetScheduleRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetSchedule(req).Return(nil)
				mocks.db.Schedule.EXPECT().Get(gomock.Any(), time.Monday).Return(schedule, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetScheduleResponse{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: true,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.GetScheduleRequest{}
				mocks.validator.EXPECT().GetSchedule(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetScheduleRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get schedule",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetSchedule(req).Return(nil)
				mocks.db.Schedule.EXPECT().Get(gomock.Any(), time.Monday).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *classroomService) (proto.Message, error) {
			return service.GetSchedule(ctx, tt.req)
		}))
	}
}
