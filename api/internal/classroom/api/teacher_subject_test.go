package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestMultiGetTeacherSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.MultiGetTeacherSubjectsRequest{
		TeacherIds: []string{"teacherid1", "teacherid2"},
	}
	teachersubjects := entity.TeacherSubjects{
		{
			TeacherID: "teacherid1",
			SubjectID: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			TeacherID: "teacherid1",
			SubjectID: 2,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			TeacherID: "teacherid2",
			SubjectID: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	subjects := entity.Subjects{
		{
			ID:         1,
			Name:       "国語",
			SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Name:       "数学",
			SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.MultiGetTeacherSubjectsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid1", "teacherid2"}
				subjectIDs := []int64{1, 2}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, subjectIDs).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: []*classroom.TeacherSubject{
						{TeacherId: "teacherid1", SubjectIds: []int64{1, 2}},
						{TeacherId: "teacherid2", SubjectIds: []int64{1}},
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
				req := &classroom.MultiGetTeacherSubjectsRequest{}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.MultiGetTeacherSubjectsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list by teacher ids",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid1", "teacherid2"}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid1", "teacherid2"}
				subjectIDs := []int64{1, 2}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, subjectIDs).Return(nil, errmock)
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
			return service.MultiGetTeacherSubjects(ctx, tt.req)
		}))
	}
}

func TestGetTeacherSubject(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.GetTeacherSubjectRequest{
		TeacherId: "teacherid",
	}
	teachersubjects := entity.TeacherSubjects{
		{
			TeacherID: "teacherid",
			SubjectID: 1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			TeacherID: "teacherid",
			SubjectID: 2,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	subjects := entity.Subjects{
		{
			ID:         1,
			Name:       "国語",
			SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Name:       "数学",
			SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetTeacherSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid"}
				subjectIDs := []int64{1, 2}
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, subjectIDs).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetTeacherSubjectResponse{
					TeacherSubject: &classroom.TeacherSubject{
						TeacherId:  "teacherid",
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
				req := &classroom.GetTeacherSubjectRequest{}
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetTeacherSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list by teacher ids",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid"}
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(nil, errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid"}
				subjectIDs := []int64{1, 2}
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, subjectIDs).Return(nil, errmock)
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
			return service.GetTeacherSubject(ctx, tt.req)
		}))
	}
}
