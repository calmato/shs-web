package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/classroom/database"
	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/internal/classroom/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.ListSubjectsRequest{
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	}
	params := &database.ListSubjectsParams{
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
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
		{
			ID:         3,
			Name:       "英語",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.ListSubjectsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListSubjects(req).Return(nil)
				mocks.db.Subject.EXPECT().List(gomock.Any(), params).Return(subjects, nil)
				mocks.db.Subject.EXPECT().Count(gomock.Any()).Return(int64(3), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.ListSubjectsResponse{
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
						{
							Id:         3,
							Name:       "英語",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
					},
					Total: 3,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.ListSubjectsRequest{}
				mocks.validator.EXPECT().ListSubjects(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.ListSubjectsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListSubjects(req).Return(nil)
				mocks.db.Subject.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.Subject.EXPECT().Count(gomock.Any()).Return(int64(3), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to count",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListSubjects(req).Return(nil)
				mocks.db.Subject.EXPECT().List(gomock.Any(), params).Return(subjects, nil)
				mocks.db.Subject.EXPECT().Count(gomock.Any()).Return(int64(0), errmock)
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
			return service.ListSubjects(ctx, tt.req)
		}))
	}
}

func TestMultiGetSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.MultiGetSubjectsRequest{
		Ids: []int64{1, 2, 3},
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
		{
			ID:         3,
			Name:       "英語",
			SchoolType: entity.SchoolTypeHighSchool,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.MultiGetSubjectsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().MultiGetSubjects(req).Return(nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, []int64{1, 2, 3}).Return(subjects, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.MultiGetSubjectsResponse{
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
						{
							Id:         3,
							Name:       "英語",
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
				req := &classroom.MultiGetSubjectsRequest{}
				mocks.validator.EXPECT().MultiGetSubjects(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.MultiGetSubjectsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().MultiGetSubjects(req).Return(nil)
				mocks.db.Subject.EXPECT().MultiGet(ctx, []int64{1, 2, 3}).Return(nil, errmock)
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
			return service.MultiGetSubjects(ctx, tt.req)
		}))
	}
}

func TestGetSubject(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &classroom.GetSubjectRequest{
		Id: 1,
	}
	subject := &entity.Subject{
		ID:         1,
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: entity.SchoolTypeHighSchool,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.GetSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Get(ctx, int64(1)).Return(subject, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.GetSubjectResponse{
					Subject: &classroom.Subject{
						Id:         1,
						Name:       "国語",
						Color:      "#F8BBD0",
						SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						CreatedAt:  now.Unix(),
						UpdatedAt:  now.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.GetSubjectRequest{}
				mocks.validator.EXPECT().GetSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.GetSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to get subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Get(ctx, int64(1)).Return(nil, errmock)
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
			return service.GetSubject(ctx, tt.req)
		}))
	}
}

func TestCreateSubject(t *testing.T) {
	t.Parallel()
	req := &classroom.CreateSubjectRequest{
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	}
	subject := &entity.Subject{
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: entity.SchoolTypeHighSchool,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.CreateSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Create(ctx, subject).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.CreateSubjectResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.CreateSubjectRequest{}
				mocks.validator.EXPECT().CreateSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.CreateSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.CreateSubjectRequest{
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
				}
				mocks.validator.EXPECT().CreateSubject(req).Return(nil)
			},
			req: &classroom.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to create subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Create(ctx, subject).Return(errmock)
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
			return service.CreateSubject(ctx, tt.req)
		}))
	}
}

func TestUpdateSubject(t *testing.T) {
	t.Parallel()
	req := &classroom.UpdateSubjectRequest{
		Id:         1,
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	}
	subject := &entity.Subject{
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: entity.SchoolTypeHighSchool,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.UpdateSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Update(ctx, int64(1), subject).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.UpdateSubjectResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpdateSubjectRequest{}
				mocks.validator.EXPECT().UpdateSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.UpdateSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.UpdateSubjectRequest{
					Id:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
				}
				mocks.validator.EXPECT().UpdateSubject(req).Return(nil)
			},
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Update(ctx, int64(1), subject).Return(errmock)
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
			return service.UpdateSubject(ctx, tt.req)
		}))
	}
}

func TestDeleteSubject(t *testing.T) {
	t.Parallel()
	req := &classroom.DeleteSubjectRequest{
		Id: 1,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *classroom.DeleteSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Delete(ctx, int64(1)).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &classroom.DeleteSubjectResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &classroom.DeleteSubjectRequest{}
				mocks.validator.EXPECT().DeleteSubject(req).Return(validation.ErrRequestValidation)
			},
			req: &classroom.DeleteSubjectRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to delete subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteSubject(req).Return(nil)
				mocks.db.Subject.EXPECT().Delete(ctx, int64(1)).Return(errmock)
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
			return service.DeleteSubject(ctx, tt.req)
		}))
	}
}
