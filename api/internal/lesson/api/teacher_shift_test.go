package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestUpsertTeacherShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.UpsertTeacherShiftsRequest{
		TeacherId:      "teacherid",
		ShiftSummaryId: 1,
		ShiftIds:       []int64{1, 2},
		Desided:        true,
	}
	teacher := &user.Teacher{Id: "teacherid"}
	shifts := entity.Shifts{
		{ID: 1, ShiftSummaryID: 1},
		{ID: 2, ShiftSummaryID: 1},
	}
	teacherSubmission := &entity.TeacherSubmission{
		TeacherID:      "teacherid",
		ShiftSummaryID: 1,
		Decided:        true,
	}
	teacherShifts := entity.TeacherShifts{
		{
			TeacherID:      "teacherid",
			ShiftID:        1,
			ShiftSummaryID: 1,
		},
		{
			TeacherID:      "teacherid",
			ShiftID:        2,
			ShiftSummaryID: 1,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.UpsertTeacherShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
				mocks.db.TeacherShift.EXPECT().Replace(ctx, teacherSubmission, teacherShifts).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.UpsertTeacherShiftsResponse{
					Submission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						Status:         lesson.TeacherSubmissionStatus_TEACHER_SUBMISSION_STATUS_SUBMITTED,
						CreatedAt:      time.Time{}.Unix(),
						UpdatedAt:      time.Time{}.Unix(),
					},
					Shifts: []*lesson.TeacherShift{
						{
							TeacherId:      "teacherid",
							ShiftSummaryId: 1,
							ShiftId:        1,
							CreatedAt:      time.Time{}.Unix(),
							UpdatedAt:      time.Time{}.Unix(),
						},
						{
							TeacherId:      "teacherid",
							ShiftSummaryId: 1,
							ShiftId:        2,
							CreatedAt:      time.Time{}.Unix(),
							UpdatedAt:      time.Time{}.Unix(),
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.UpsertTeacherShiftsRequest{}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.UpsertTeacherShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(nil, errmock)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to shifts length is unmatch",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				shifts := entity.Shifts{}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to replace teacher shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
				mocks.db.TeacherShift.EXPECT().Replace(ctx, teacherSubmission, teacherShifts).Return(errmock)
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
			return service.UpsertTeacherShifts(ctx, tt.req)
		}))
	}
}