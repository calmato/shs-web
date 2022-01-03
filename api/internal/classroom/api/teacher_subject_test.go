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

func TestMultiGetTeacherSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.MultiGetTeacherSubjectsRequest{
		TeacherIds: []string{"teacherid1"},
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
		req    *classroom.MultiGetTeacherSubjectsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid1"}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: []*classroom.TeacherSubject{
						{TeacherId: "teacherid1", SubjectIds: []int64{1, 2}},
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
				teacherIDs := []string{"teacherid1"}
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
				teacherIDs := []string{"teacherid1"}
				mocks.validator.EXPECT().MultiGetTeacherSubjects(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
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
		req    *classroom.GetTeacherSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				teacherIDs := []string{"teacherid"}
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, gomock.Any()).Return(subjects, nil)
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
				mocks.validator.EXPECT().GetTeacherSubject(req).Return(nil)
				mocks.db.TeacherSubject.EXPECT().ListByTeacherIDs(ctx, teacherIDs).Return(teachersubjects, nil)
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
			return service.GetTeacherSubject(ctx, tt.req)
		}))
	}
}

func TestUpsertTeacherSubject(t *testing.T) {
	t.Parallel()

	now := jst.Now()
	req := &classroom.UpsertTeacherSubjectRequest{
		TeacherId:  "teacherid",
		SubjectIds: []int64{1, 2},
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	}
	teacher := &user.Teacher{Id: "teacherid"}
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
	teachersubjects := entity.TeacherSubjects{
		{
			TeacherID: "teacherid",
			SubjectID: 1,
		},
		{
			TeacherID: "teacherid",
			SubjectID: 2,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.UpsertTeacherSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetTeacherRequest{Id: "teacherid"}
				out := &user.GetTeacherResponse{Teacher: teacher}
				schoolType := entity.SchoolTypeHighSchool
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
				mocks.db.TeacherSubject.EXPECT().Replace(ctx, schoolType, teachersubjects).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.UpsertTeacherSubjectResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpsertTeacherSubjectRequest{}
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.UpsertTeacherSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpsertTeacherSubjectRequest{}
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
			},
			req: &classroom.UpsertTeacherSubjectRequest{
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetTeacherRequest{Id: "teacherid"}
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(nil, errmock)
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
				in := &user.GetTeacherRequest{Id: "teacherid"}
				out := &user.GetTeacherResponse{Teacher: teacher}
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
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
				req := &classroom.UpsertTeacherSubjectRequest{
					TeacherId:  "teacherid",
					SubjectIds: []int64{1, 2},
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				in := &user.GetTeacherRequest{Id: "teacherid"}
				out := &user.GetTeacherResponse{Teacher: teacher}
				subjects := entity.Subjects{}
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
			},
			req: &classroom.UpsertTeacherSubjectRequest{
				TeacherId:  "teacherid",
				SubjectIds: []int64{1, 2},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to replace teacher subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				in := &user.GetTeacherRequest{Id: "teacherid"}
				out := &user.GetTeacherResponse{Teacher: teacher}
				schoolType := entity.SchoolTypeHighSchool
				mocks.validator.EXPECT().UpsertTeacherSubject(req).Return(nil)
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
				mocks.db.Subject.EXPECT().MultiGet(gomock.Any(), gomock.Any()).Return(subjects, nil)
				mocks.db.TeacherSubject.EXPECT().Replace(ctx, schoolType, teachersubjects).Return(errmock)
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
			return service.UpsertTeacherSubject(ctx, tt.req)
		}))
	}
}
