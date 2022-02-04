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
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"gorm.io/datatypes"
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
			SuggestedLessons: entity.SuggestedLessons{
				{SubjectID: 1, Total: 4},
				{SubjectID: 2, Total: 4},
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			StudentID:      "studentid",
			ShiftSummaryID: 2,
			Decided:        false,
			SuggestedLessons: entity.SuggestedLessons{
				{SubjectID: 1, Total: 8},
			},
			CreatedAt: now,
			UpdatedAt: now,
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
							SuggestedLessons: []*lesson.SuggestedLesson{
								{SubjectId: 1, Total: 4},
								{SubjectId: 2, Total: 4},
							},
							CreatedAt: now.Unix(),
							UpdatedAt: now.Unix(),
						},
						{
							StudentId:      "studentid",
							ShiftSummaryId: 2,
							Decided:        false,
							SuggestedLessons: []*lesson.SuggestedLesson{
								{SubjectId: 1, Total: 8},
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
			SuggestedLessons: entity.SuggestedLessons{
				{SubjectID: 1, Total: 4},
				{SubjectID: 2, Total: 4},
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			StudentID:        "studentid2",
			ShiftSummaryID:   1,
			Decided:          false,
			SuggestedLessons: entity.SuggestedLessons{},
			CreatedAt:        now,
			UpdatedAt:        now,
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
							SuggestedLessons: []*lesson.SuggestedLesson{
								{SubjectId: 1, Total: 4},
								{SubjectId: 2, Total: 4},
							},
							CreatedAt: now.Unix(),
							UpdatedAt: now.Unix(),
						},
						{
							StudentId:        "studentid2",
							ShiftSummaryId:   1,
							Decided:          false,
							SuggestedLessons: []*lesson.SuggestedLesson{},
							CreatedAt:        now.Unix(),
							UpdatedAt:        now.Unix(),
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
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
			{SubjectID: 2, Total: 4},
		},
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
						SuggestedLessons: []*lesson.SuggestedLesson{
							{SubjectId: 1, Total: 4},
							{SubjectId: 2, Total: 4},
						},
						CreatedAt: time.Time{}.Unix(),
						UpdatedAt: time.Time{}.Unix(),
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
			name: "success to submission is not found",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShifts(req).Return(nil)
				mocks.db.StudentSubmission.EXPECT().Get(gomock.Any(), "studentid", int64(1)).Return(nil, database.ErrNotFound)
				mocks.db.StudentShift.EXPECT().ListByShiftSummaryID(gomock.Any(), []string{"studentid"}, int64(1)).Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.GetStudentShiftsResponse{
					Submission: &lesson.StudentSubmission{
						StudentId:        "",
						ShiftSummaryId:   0,
						Decided:          false,
						SuggestedLessons: []*lesson.SuggestedLesson{},
						CreatedAt:        time.Time{}.Unix(),
						UpdatedAt:        time.Time{}.Unix(),
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
		Lessons: []*lesson.StudentSuggestedLesson{
			{SubjectId: 1, Total: 4},
			{SubjectId: 2, Total: 4},
		},
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
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
			{SubjectID: 2, Total: 4},
		},
	}
	_ = studentSubmission.FillJSON()
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
	studentSubject := &classroom.StudentSubject{
		StudentId:  "studentid",
		SubjectIds: []int64{1, 2},
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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
						SuggestedLessons: []*lesson.SuggestedLesson{
							{SubjectId: 1, Total: 4},
							{SubjectId: 2, Total: 4},
						},
						CreatedAt: time.Time{}.Unix(),
						UpdatedAt: time.Time{}.Unix(),
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get student subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(nil, errmock)
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				shifts := entity.Shifts{}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				summary := &entity.ShiftSummary{
					ID:     1,
					Status: entity.ShiftStatusFinished,
				}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: &classroom.StudentSubject{}}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
				mocks.db.ShiftSummary.EXPECT().Get(gomock.Any(), int64(1)).Return(summary, nil)
				mocks.db.Shift.EXPECT().MultiGet(gomock.Any(), []int64{1, 2}, "id").Return(shifts, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to replace student shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIn := &user.GetStudentRequest{Id: "studentid"}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectIn := &classroom.GetStudentSubjectRequest{StudentId: "studentid"}
				subjectOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject}
				mocks.validator.EXPECT().UpsertStudentShifts(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(ctx, subjectIn).Return(subjectOut, nil)
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

func TestGetStudentShiftTempalte(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	req := &lesson.GetStudentShiftTemplateRequest{
		StudentId: "studentid",
	}
	template := &entity.StudentShiftTemplate{
		StudentID: "studentid",
		Schedules: entity.ShiftSchedules{
			{
				Weekday: time.Sunday,
				Lessons: entity.LessonSchedules{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
				},
			},
		},
		SchedulesJSON: datatypes.JSON([]byte(`[{"weekday":0,"lessons":[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"2000"}]}]`)),
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
		},
		SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4}]`)),
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.GetStudentShiftTemplateRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShiftTemplate(req).Return(nil)
				mocks.db.StudentShiftTemplate.EXPECT().Get(ctx, "studentid").Return(template, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.GetStudentShiftTemplateResponse{
					Template: &lesson.StudentShiftTemplate{
						StudentId: "studentid",
						Schedules: []*lesson.ShiftSchedule{
							{
								Weekday: int32(time.Sunday),
								Lessons: []*lesson.LessonSchedule{
									{StartTime: "1700", EndTime: "1830"},
									{StartTime: "1830", EndTime: "2000"},
								},
							},
						},
						SuggestedLessons: []*lesson.SuggestedLesson{
							{SubjectId: 1, Total: 4},
						},
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
		},
		{
			name: "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.GetStudentShiftTemplateRequest{}
				mocks.validator.EXPECT().GetStudentShiftTemplate(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.GetStudentShiftTemplateRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get student shift template",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudentShiftTemplate(req).Return(nil)
				mocks.db.StudentShiftTemplate.EXPECT().Get(ctx, "studentid").Return(nil, errmock)
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
			return service.GetStudentShiftTemplate(ctx, tt.req)
		}))
	}
}

func TestUpsertStudentShiftTempalte(t *testing.T) {
	t.Parallel()
	req := &lesson.UpsertStudentShiftTemplateRequest{
		StudentId: "studentid",
		Template: &lesson.StudentShiftTemplateToUpsert{
			Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
				{
					Weekday: int32(time.Sunday),
					Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
					},
				},
			},
			SuggestedLessons: []*lesson.StudentSuggestedLesson{
				{SubjectId: 1, Total: 4},
			},
		},
	}
	student := &user.Student{Id: "studentid"}
	template := &entity.StudentShiftTemplate{
		StudentID: "studentid",
		Schedules: entity.ShiftSchedules{
			{
				Weekday: time.Sunday,
				Lessons: entity.LessonSchedules{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
				},
			},
		},
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *lesson.UpsertStudentShiftTemplateRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShiftTemplate(req).Return(nil)
				mocks.user.EXPECT().GetStudent(ctx, in).Return(out, nil)
				mocks.db.StudentShiftTemplate.EXPECT().Upsert(ctx, "studentid", template).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &lesson.UpsertStudentShiftTemplateResponse{},
			},
		},
		{
			name: "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &lesson.UpsertStudentShiftTemplateRequest{}
				mocks.validator.EXPECT().UpsertStudentShiftTemplate(req).Return(validation.ErrRequestValidation)
			},
			req: &lesson.UpsertStudentShiftTemplateRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				mocks.validator.EXPECT().UpsertStudentShiftTemplate(req).Return(nil)
				mocks.user.EXPECT().GetStudent(ctx, in).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to get student shift template",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentShiftTemplate(req).Return(nil)
				mocks.user.EXPECT().GetStudent(ctx, in).Return(out, nil)
				mocks.db.StudentShiftTemplate.EXPECT().Upsert(ctx, "studentid", template).Return(errmock)
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
			return service.UpsertStudentShiftTemplate(ctx, tt.req)
		}))
	}
}
