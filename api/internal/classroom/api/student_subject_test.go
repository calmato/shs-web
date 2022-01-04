package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestMultiGetStudentSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.MultiGetStudentSubjectsRequest{
		StudentIds: []string{"studentid1"},
	}
	studentsubjects := entity.StudentSubjects{
		{
			StudentID: "studentid1",
			SubjectID: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			StudentID: "studentid1",
			SubjectID: 2,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	subjects := entity.Subjects{
		{
			ID:         1,
			Name:       "国語",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Name:       "数学",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.MultiGetStudentSubjectsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid1"}
				mocks.validator.EXPECT().MultiGetStudentSubjects(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(studentsubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: []*classroom.StudentSubject{
						{StudentId: "studentid1", SubjectIds: []int64{1, 2}},
					},
					Subjects: []*classroom.Subject{
						{
							Id:         1,
							Name:       "国語",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
						{
							Id:         2,
							Name:       "数学",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.MultiGetStudentSubjectsRequest{}
				mocks.validator.EXPECT().MultiGetStudentSubjects(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.MultiGetStudentSubjectsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list by student ids",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid1"}
				mocks.validator.EXPECT().MultiGetStudentSubjects(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid1"}
				mocks.validator.EXPECT().MultiGetStudentSubjects(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(studentsubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(nil, errmock)
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
			return service.MultiGetStudentSubjects(ctx, tt.req)
		}))
	}
}

func TestGetStudentSubject(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.GetStudentSubjectRequest{
		StudentId: "studentid",
	}
	studentsubjects := entity.StudentSubjects{
		{
			StudentID: "studentid",
			SubjectID: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			StudentID: "studentid",
			SubjectID: 2,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	subjects := entity.Subjects{
		{
			ID:         1,
			Name:       "国語",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Name:       "数学",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetStudentSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid"}
				mocks.validator.EXPECT().GetStudentSubject(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(studentsubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetStudentSubjectResponse{
					StudentSubject: &classroom.StudentSubject{
						StudentId:  "studentid",
						SubjectIds: []int64{1, 2},
					},
					Subjects: []*classroom.Subject{
						{
							Id:         1,
							Name:       "国語",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
						{
							Id:         2,
							Name:       "数学",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.GetStudentSubjectRequest{}
				mocks.validator.EXPECT().GetStudentSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetStudentSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list by student ids",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid"}
				mocks.validator.EXPECT().GetStudentSubject(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				studentIDs := []string{"studentid"}
				mocks.validator.EXPECT().GetStudentSubject(req).Return(nil)
				mocks.db.StudentSubject.EXPECT().ListByStudentIDs(ctx, studentIDs).Return(studentsubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(nil, errmock)
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
			return service.GetStudentSubject(ctx, tt.req)
		}))
	}
}

func TestUpsertStudentSubject(t *testing.T) {
	t.Parallel()

	now := jst.Now()
	req := &classroom.UpsertStudentSubjectRequest{
		StudentId:  "studentid",
		SubjectIds: []int64{1, 2},
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	}
	student := &user.Student{Id: "studentid"}
	subjects := entity.Subjects{
		{
			ID:         1,
			Name:       "国語",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Name:       "数学",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}
	studentsubjects := entity.StudentSubjects{
		{
			StudentID: "studentid",
			SubjectID: 1,
		},
		{
			StudentID: "studentid",
			SubjectID: 2,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.UpsertStudentSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				schoolType := entity.SchoolTypeHighSchool
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
				mocks.db.StudentSubject.EXPECT().Replace(ctx, schoolType, studentsubjects).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.UpsertStudentSubjectResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpsertStudentSubjectRequest{}
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.UpsertStudentSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpsertStudentSubjectRequest{}
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
			},
			req: &classroom.UpsertStudentSubjectRequest{
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(nil, errmock)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to unmatch subject ids",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpsertStudentSubjectRequest{
					StudentId:  "studentid",
					SubjectIds: []int64{1, 2},
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				subjects := entity.Subjects{}
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
			},
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "studentid",
				SubjectIds: []int64{1, 2},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to replace student subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetStudentRequest{Id: "studentid"}
				out := &user.GetStudentResponse{Student: student}
				schoolType := entity.SchoolTypeHighSchool
				mocks.validator.EXPECT().UpsertStudentSubject(req).Return(nil)
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
				mocks.db.StudentSubject.EXPECT().Replace(ctx, schoolType, studentsubjects).Return(errmock)
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
			return service.UpsertStudentSubject(ctx, tt.req)
		}))
	}
}
