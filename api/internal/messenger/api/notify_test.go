package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/messenger/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestNotifyLessonDecided(t *testing.T) {
	t.Parallel()

	now := jst.Now()
	req := &messenger.NotifyLessonDecidedRequest{
		ShiftSummaryId: 1,
	}
	summary := &lesson.ShiftSummary{
		Id:        1,
		YearMonth: 202202,
		Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
		OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
		EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
	lessons := []*lesson.Lesson{
		{
			Id:             1,
			ShiftSummaryId: 1,
			ShiftId:        1,
			SubjectId:      1,
			RoomId:         1,
			TeacherId:      "teacherid",
			StudentId:      "studentid",
			Notes:          "",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *messenger.NotifyLessonDecidedRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsResponse{Lessons: lessons}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.publisher.EXPECT().Publish(ctx, gomock.Any()).Return("success", nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &messenger.NotifyLessonDecidedResponse{},
			},
		},
		{
			name: "success to lessons is length 0",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsResponse{Lessons: []*lesson.Lesson{}}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &messenger.NotifyLessonDecidedResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &messenger.NotifyLessonDecidedRequest{}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(validation.ErrRequestValidation)
			},
			req: &messenger.NotifyLessonDecidedRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsResponse{Lessons: lessons}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to list lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to parse year month",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summary := &lesson.ShiftSummary{
					Id:        1,
					YearMonth: 202200,
					Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsResponse{Lessons: lessons}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to publish",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				lessonsIn := &lesson.ListLessonsRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsResponse{Lessons: lessons}
				mocks.validator.EXPECT().NotifyLessonDecided(req).Return(nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.publisher.EXPECT().Publish(ctx, gomock.Any()).Return("", errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *messengerService) (proto.Message, error) {
			return service.NotifyLessonDecided(ctx, tt.req)
		}))
	}
}
