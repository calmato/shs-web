package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestCreateShifts(t *testing.T) {
	t.Parallel()

	req := &lesson.CreateShiftsRequest{
		YearMonth:   202202,
		OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
		EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, 0).Unix(),
		ClosedDates: []string{"20220208", "20220221", "20220222"},
	}
	schedules := []*classroom.Schedule{
		{Weekday: int32(time.Sunday), IsClosed: true},
		{
			Weekday:  int32(time.Monday),
			IsClosed: false,
			Lessons: []*classroom.Schedule_Lesson{
				{StartTime: "1530", EndTime: "1700"},
			},
		},
		{
			Weekday:  int32(time.Tuesday),
			IsClosed: false,
			Lessons: []*classroom.Schedule_Lesson{
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
			},
		},
		{Weekday: int32(time.Wednesday), IsClosed: true},
		{Weekday: int32(time.Thursday), IsClosed: true},
		{Weekday: int32(time.Friday), IsClosed: true},
		{Weekday: int32(time.Saturday), IsClosed: true},
	}
	summary := &entity.ShiftSummary{
		YearMonth: 202202,
		OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
		EndAt:     jst.Date(2022, 1, 14, 23, 59, 59, 0),
	}
	shifts := entity.Shifts{
		{Date: jst.Date(2022, 2, 1, 0, 0, 0, 0), StartTime: "1700", EndTime: "1830"},
		{Date: jst.Date(2022, 2, 1, 0, 0, 0, 0), StartTime: "1830", EndTime: "2000"},
		{Date: jst.Date(2022, 2, 7, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
		{Date: jst.Date(2022, 2, 14, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
		{Date: jst.Date(2022, 2, 15, 0, 0, 0, 0), StartTime: "1700", EndTime: "1830"},
		{Date: jst.Date(2022, 2, 15, 0, 0, 0, 0), StartTime: "1830", EndTime: "2000"},
		{Date: jst.Date(2022, 2, 28, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.CreateShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &classroom.ListSchedulesRequest{}
				out := &classroom.ListSchedulesResponse{Schedules: schedules}
				mocks.validator.EXPECT().CreateShifts(req).Return(nil)
				mocks.classroom.EXPECT().ListSchedules(ctx, in).Return(out, nil)
				mocks.db.Shift.EXPECT().
					MultipleCreate(ctx, summary, gomock.Any()).
					DoAndReturn(func(ctx context.Context, summary *entity.ShiftSummary, targets entity.Shifts) error {
						if len(targets) != len(shifts) {
							return errmock
						}
						return nil
					})
			},
			req: req,
			expect: &testResponse{
				code: codes.OK, // ignore: レスポンスが大きすぎるので、ステータスチェックのみ
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.CreateShiftsRequest{}
				mocks.validator.EXPECT().CreateShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.CreateShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list classroom schedules",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &classroom.ListSchedulesRequest{}
				mocks.validator.EXPECT().CreateShifts(req).Return(nil)
				mocks.classroom.EXPECT().ListSchedules(ctx, in).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to new shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.CreateShiftsRequest{
					YearMonth:   202200,
					OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, 0).Unix(),
					ClosedDates: []string{"20220208", "20220221", "20220222"},
				}
				in := &classroom.ListSchedulesRequest{}
				out := &classroom.ListSchedulesResponse{Schedules: schedules}
				mocks.validator.EXPECT().CreateShifts(req).Return(nil)
				mocks.classroom.EXPECT().ListSchedules(ctx, in).Return(out, nil)
			},
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202200,
				OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220208", "20220221", "20220222"},
			},
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multiple create",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &classroom.ListSchedulesRequest{}
				out := &classroom.ListSchedulesResponse{Schedules: schedules}
				mocks.validator.EXPECT().CreateShifts(req).Return(nil)
				mocks.classroom.EXPECT().ListSchedules(ctx, in).Return(out, nil)
				mocks.db.Shift.EXPECT().MultipleCreate(ctx, summary, gomock.Any()).Return(errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *lessonService) (proto.Message, error) {
			return service.CreateShifts(ctx, tt.req)
		}))
	}
}
