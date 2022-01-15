package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/internal/lesson/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListStudentSubmissionsByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	req := &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
		StudentId:       "studentid",
		ShiftSummaryIds: []int64{1, 2},
	}
	submissions := entity.StudentSubmissions{
		{
			StudentID:      "studentid",
			ShiftSummaryID: 1,
			Decided:        true,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			StudentID:      "studentid",
			ShiftSummaryID: 2,
			Decided:        false,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentSubmissionsByShiftSummaryIDs(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().ListByShiftSummaryIDs(ctx, "studentid", []int64{1, 2}).Return(submissions, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListStudentSubmissionsByShiftSummaryIDsResponse{
					Submissions: []*lesson.StudentSubmission{
						{
							StudentId:      "studentid",
							ShiftSummaryId: 1,
							Decided:        true,
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
						{
							StudentId:      "studentid",
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
				mocks.validator.EXPECT().ListStudentSubmissionsByShiftSummaryIDs(req).Return(validation.ErrRequestValidation)
			},
			req: req,
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list student submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentSubmissionsByShiftSummaryIDs(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().ListByShiftSummaryIDs(ctx, "studentid", []int64{1, 2}).Return(nil, errmock)
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
			return service.ListStudentSubmissionsByShiftSummaryIDs(ctx, tt.req)
		}))
	}
}

func TestListStudentSubmissionsByStudentIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	req := &lesson.ListStudentSubmissionsByStudentIDsRequest{
		StudentIds:     []string{"studentid"},
		ShiftSummaryId: 1,
	}
	submissions := entity.StudentSubmissions{
		{
			StudentID:      "studentid1",
			ShiftSummaryID: 1,
			Decided:        true,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			StudentID:      "studentid2",
			ShiftSummaryID: 1,
			Decided:        false,
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListStudentSubmissionsByStudentIDsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentSubmissionsByStudentIDs(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().ListByStudentIDs(ctx, []string{"studentid"}, int64(1)).Return(submissions, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListStudentSubmissionsByStudentIDsResponse{
					Submissions: []*lesson.StudentSubmission{
						{
							StudentId:      "studentid1",
							ShiftSummaryId: 1,
							Decided:        true,
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
						{
							StudentId:      "studentid2",
							ShiftSummaryId: 1,
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
				req := &lesson.ListStudentSubmissionsByStudentIDsRequest{}
				mocks.validator.EXPECT().ListStudentSubmissionsByStudentIDs(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListStudentSubmissionsByStudentIDsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list student submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentSubmissionsByStudentIDs(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().ListByStudentIDs(ctx, []string{"studentid"}, int64(1)).Return(nil, errmock)
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
			return service.ListStudentSubmissionsByStudentIDs(ctx, tt.req)
		}))
	}
}

func TestListStudentShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.ListStudentShiftsRequest{
		StudentIds:     []string{"studentid1", "studentid2"},
		ShiftSummaryId: 1,
	}
	shifts := entity.StudentShifts{
		{
			StudentID:      "studentid",
			ShiftID:        1,
			ShiftSummaryID: 1,
		},
		{
			StudentID:      "studentid",
			ShiftID:        2,
			ShiftSummaryID: 1,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.ListStudentShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentShifts(req).Return(nil)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(ctx, []string{"studentid1", "studentid2"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.ListStudentShiftsResponse{
					Shifts: []*lesson.StudentShift{
						{
							StudentId:      "studentid",
							ShiftSummaryId: 1,
							ShiftId:        1,
							CreatedAt:      time.Time{}.Unix(),
							UpdatedAt:      time.Time{}.Unix(),
						},
						{
							StudentId:      "studentid",
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
				req := &lesson.ListStudentShiftsRequest{}
				mocks.validator.EXPECT().ListStudentShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.ListStudentShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudentShifts(req).Return(nil)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(ctx, []string{"studentid1", "studentid2"}, int64(1)).Return(nil, errmock)
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
			return service.ListStudentShifts(ctx, tt.req)
		}))
	}
}

func TestGetStudentShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.GetStudentShiftsRequest{
		StudentId:      "studentid",
		ShiftSummaryId: 1,
	}
	submission := &entity.StudentSubmission{
		StudentID:      "studentid",
		ShiftSummaryID: 1,
		Decided:        true,
	}
	shifts := entity.StudentShifts{
		{
			StudentID:      "studentid",
			ShiftID:        1,
			ShiftSummaryID: 1,
		},
		{
			StudentID:      "studentid",
			ShiftID:        2,
			ShiftSummaryID: 1,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.GetStudentShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShifts(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().Get(gomock.Any(), "studentid", int64(1)).Return(submission, nil)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"studentid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.GetStudentShiftsResponse{
					Submission: &lesson.StudentSubmission{
						StudentId:      "studentid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      time.Time{}.Unix(),
						UpdatedAt:      time.Time{}.Unix(),
					},
					Shifts: []*lesson.StudentShift{
						{
							StudentId:      "studentid",
							ShiftSummaryId: 1,
							ShiftId:        1,
							CreatedAt:      time.Time{}.Unix(),
							UpdatedAt:      time.Time{}.Unix(),
						},
						{
							StudentId:      "studentid",
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
				req := &lesson.GetStudentShiftsRequest{}
				mocks.validator.EXPECT().GetStudentShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.GetStudentShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get submission",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShifts(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().Get(gomock.Any(), "studentid", int64(1)).Return(nil, errmock)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"studentid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShifts(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().Get(gomock.Any(), "studentid", int64(1)).Return(submission, nil)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"studentid"}, int64(1)).Return(nil, errmock)
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
			return service.GetStudentShifts(ctx, tt.req)
		}))
	}
}

func TestUpsertStudentShifts(t *testing.T) {
	t.Parallel()
	req := &lesson.UpsertStudentShiftsRequest{
		StudentId:      "studentid",
		ShiftSummaryId: 1,
		ShiftIds:       []int64{1, 2},
		Decided:        true,
	}
	student := &user.Student{Id: "studentid"}
	summary := &entity.ShiftSummary{
		ID:     1,
		Status: entity.ShiftStatusAccepting,
	}
	shifts := entity.Shifts{
		{ID: 1, ShiftSummaryID: 1},
		{ID: 2, ShiftSummaryID: 1},
	}
	studentSubmission := &entity.StudentSubmission{
		StudentID:      "studentid",
		ShiftSummaryID: 1,
		Decided:        true,
	}
	studentShifts := entity.StudentShifts{
		{
			StudentID:      "studentid",
			ShiftID:        1,
			ShiftSummaryID: 1,
		},
		{
			StudentID:      "studentid",
			ShiftID:        2,
			ShiftSummaryID: 1,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.UpsertStudentShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
				mocks.db.StudentShift.EXPECT().Replace(ctx, studentSubmission, studentShifts).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.UpsertStudentShiftsResponse{
					Submission: &lesson.StudentSubmission{
						StudentId:      "studentid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      time.Time{}.Unix(),
						UpdatedAt:      time.Time{}.Unix(),
					},
					Shifts: []*lesson.StudentShift{
						{
							StudentId:      "studentid",
							ShiftSummaryId: 1,
							ShiftId:        1,
							CreatedAt:      time.Time{}.Unix(),
							UpdatedAt:      time.Time{}.Unix(),
						},
						{
							StudentId:      "studentid",
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
				req := &lesson.UpsertStudentShiftsRequest{}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.UpsertStudentShiftsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
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
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
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
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
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
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				shifts := entity.Shifts{}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
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
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				summary := &entity.ShiftSummary{
					ID:     1,
					Status: entity.ShiftStatusFinished,
				}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.FailedPrecondition,
			},
		},
		{
			name: "failed to replace student shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
				mocks.db.StudentShift.EXPECT().Replace(ctx, studentSubmission, studentShifts).Return(errmock)
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
			return service.UpsertStudentShifts(ctx, tt.req)
		}))
	}
}
