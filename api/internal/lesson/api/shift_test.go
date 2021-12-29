package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListShiftSummaries(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &lesson.ListShiftSummariesRequest{
		Status: lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
		Limit:  30,
		Offset: 0,
	}
	params := &database.ListShiftSummariesParams{
		Status: entity.ShiftStatusAccepting,
		Limit:  30,
		Offset: 0,
	}
	summaries := entity.ShiftSummaries{
		{
			ID:        1,
			Status:    entity.ShiftStatusAccepting,
			OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
			EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListShiftSummariesRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListShiftSummaries(req).Return(nil)
				mocks.db.ShiftSummary.EXPECT().List(gomock.Any(), params).Return(summaries, nil)
				mocks.db.ShiftSummary.EXPECT().Count(gomock.Any()).Return(int64(1), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListShiftSummariesResponse{
					Summaries: []*lesson.ShiftSummary{
						{
							Id:        1,
							Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
							OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
							EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
							CreatedAt: now.Unix(),
							UpdatedAt: now.Unix(),
						},
					},
					Total: 1,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.ListShiftSummariesRequest{}
				mocks.validator.EXPECT().ListShiftSummaries(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListShiftSummariesRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list shift summaries",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListShiftSummaries(req).Return(nil)
				mocks.db.ShiftSummary.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.ShiftSummary.EXPECT().Count(gomock.Any()).Return(int64(1), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to count",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListShiftSummaries(req).Return(nil)
				mocks.db.ShiftSummary.EXPECT().List(gomock.Any(), params).Return(summaries, nil)
				mocks.db.ShiftSummary.EXPECT().Count(gomock.Any()).Return(int64(0), errmock)
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
			return service.ListShiftSummaries(ctx, tt.req)
		}))
	}
}

func TestListShifts(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &lesson.ListShiftsRequest{
		ShiftSummaryId: 1,
	}
	shifts := entity.Shifts{
		{
			ID:             1,
			ShiftSummaryID: 1,
			Date:           jst.Date(2022, 2, 1, 0, 0, 0, 0),
			StartTime:      "1700",
			EndTime:        "1830",
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             2,
			ShiftSummaryID: 1,
			Date:           jst.Date(2022, 2, 1, 0, 0, 0, 0),
			StartTime:      "1830",
			EndTime:        "2000",
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListShifts(req).Return(nil)
				mocks.db.Shift.EXPECT().ListBySummaryID(ctx, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListShiftsResponse{
					Shifts: []*lesson.Shift{
						{
							Id:             1,
							ShiftSummaryId: 1,
							Date:           "20220201",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
						{
							Id:             2,
							ShiftSummaryId: 1,
							Date:           "20220201",
							StartTime:      "1830",
							EndTime:        "2000",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.ListShiftsRequest{}
				mocks.validator.EXPECT().ListShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListShifts(req).Return(nil)
				mocks.db.Shift.EXPECT().ListBySummaryID(ctx, int64(1)).Return(nil, errmock)
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
			return service.ListShifts(ctx, tt.req)
		}))
	}
}

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
