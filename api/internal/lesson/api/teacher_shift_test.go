package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListTeacherSubmissionsByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	req := &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
		TeacherId:       "teacherid",
		ShiftSummaryIds: []int64{1, 2},
	}
	submissions := entity.TeacherSubmissions{
		{
			TeacherID:      "teacherid",
			ShiftSummaryID: 1,
			Decided:        true,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			TeacherID:      "teacherid",
			ShiftSummaryID: 2,
			Decided:        false,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeacherSubmissionsByShiftSummaryIDs(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().ListByShiftSummaryIDs(ctx, "teacherid", []int64{1, 2}).Return(submissions, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse{
					Submissions: []*lesson.TeacherSubmission{
						{
							TeacherId:      "teacherid",
							ShiftSummaryId: 1,
							Decided:        true,
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
						{
							TeacherId:      "teacherid",
							ShiftSummaryId: 2,
							Decided:        false,
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
				mocks.validator.EXPECT().ListTeacherSubmissionsByShiftSummaryIDs(req).Return(validation.ErrRequestValidation)
			},
			req: req,
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list teacher submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeacherSubmissionsByShiftSummaryIDs(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().ListByShiftSummaryIDs(ctx, "teacherid", []int64{1, 2}).Return(nil, errmock)
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
			return service.ListTeacherSubmissionsByShiftSummaryIDs(ctx, tt.req)
		}))
	}
}

func TestListTeacherShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.ListTeacherShiftsRequest{
		TeacherIds:     []string{"teacherid1", "teacherid2"},
		ShiftSummaryId: 1,
	}
	shifts := entity.TeacherShifts{
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
		req    *lesson.ListTeacherShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeacherShifts(req).Return(nil)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(ctx, []string{"teacherid1", "teacherid2"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListTeacherShiftsResponse{
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
			name: "invliad argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.ListTeacherShiftsRequest{}
				mocks.validator.EXPECT().ListTeacherShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListTeacherShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeacherShifts(req).Return(nil)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(ctx, []string{"teacherid1", "teacherid2"}, int64(1)).Return(nil, errmock)
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
			return service.ListTeacherShifts(ctx, tt.req)
		}))
	}
}

func TestGetTeacherShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.GetTeacherShiftsRequest{
		TeacherId:      "teacherid",
		ShiftSummaryId: 1,
	}
	submission := &entity.TeacherSubmission{
		TeacherID:      "teacherid",
		ShiftSummaryID: 1,
		Decided:        true,
	}
	shifts := entity.TeacherShifts{
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
		req    *lesson.GetTeacherShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacherShifts(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().Get(gomock.Any(), "teacherid", int64(1)).Return(submission, nil)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"teacherid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.GetTeacherShiftsResponse{
					Submission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
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
			name: "success to submission is not found",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacherShifts(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().Get(gomock.Any(), "teacherid", int64(1)).Return(nil, database.ErrNotFound)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"teacherid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.GetTeacherShiftsResponse{
					Submission: &lesson.TeacherSubmission{
						TeacherId:      "",
						ShiftSummaryId: 0,
						Decided:        false,
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
			name: "invliad argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.GetTeacherShiftsRequest{}
				mocks.validator.EXPECT().GetTeacherShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.GetTeacherShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get submission",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacherShifts(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().Get(gomock.Any(), "teacherid", int64(1)).Return(nil, errmock)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"teacherid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacherShifts(req).Return(nil)
				mocks.db.TeacherSubmission.EXPECT().Get(gomock.Any(), "teacherid", int64(1)).Return(submission, nil)
				mocks.db.TeacherShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"teacherid"}, int64(1)).Return(nil, errmock)
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
			return service.GetTeacherShifts(ctx, tt.req)
		}))
	}
}

func TestUpsertTeacherShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.UpsertTeacherShiftsRequest{
		TeacherId:      "teacherid",
		ShiftSummaryId: 1,
		ShiftIds:       []int64{1, 2},
		Decided:        true,
	}
	teacher := &user.Teacher{Id: "teacherid"}
	summary := &entity.ShiftSummary{
		ID:     1,
		Status: entity.ShiftStatusAccepting,
	}
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
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
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
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(nil, errmock)
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
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
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
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to outside of shift submission",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				summary := &entity.ShiftSummary{
					ID:     1,
					Status: entity.ShiftStatusFinished,
				}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.FailedPrecondition,
			},
		},
		{
			name: "failed to replace teacher shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIn := &user.GetTeacherRequest{Id: "teacherid"}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherShifts(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
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
