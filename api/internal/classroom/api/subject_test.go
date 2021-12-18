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
		Limit:  30,
		Offset: 0,
	}
	params := &database.ListSubjectsParams{
		Limit:  30,
		Offset: 0,
	}
	subjects := entity.Subjects{
		{ID: 1, Name: "国語", CreatedAt: now, UpdatedAt: now},
		{ID: 2, Name: "数学", CreatedAt: now, UpdatedAt: now},
		{ID: 3, Name: "英語", CreatedAt: now, UpdatedAt: now},
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
						{Id: 1, Name: "国語", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
						{Id: 2, Name: "数学", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
						{Id: 3, Name: "英語", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
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
		{ID: 1, Name: "国語", CreatedAt: now, UpdatedAt: now},
		{ID: 2, Name: "数学", CreatedAt: now, UpdatedAt: now},
		{ID: 3, Name: "英語", CreatedAt: now, UpdatedAt: now},
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
						{Id: 1, Name: "国語", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
						{Id: 2, Name: "数学", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
						{Id: 3, Name: "英語", CreatedAt: now.Unix(), UpdatedAt: now.Unix()},
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
		ID:        1,
		Name:      "国語",
		Color:     "#F8BBD0",
		CreatedAt: now,
		UpdatedAt: now,
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
						Id:        1,
						Name:      "国語",
						Color:     "#F8BBD0",
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
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
