package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestCreateTeacher(t *testing.T) {
	t.Parallel()

	req := &user.CreateTeacherRequest{
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          "teacher-test001@calmato.jp",
		Password:      "12345678",
		Role:          0,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.CreateTeacherRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK, // TODO: uuid周りの関係でBodyは無視で..
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.CreateTeacherRequest{}
				mocks.validator.EXPECT().CreateTeacher(req).Return(validation.ErrRequestValidation)
			},
			req: &user.CreateTeacherRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to create teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Create(ctx, gomock.Any()).Return(errmock)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, testGRPC(tt.setup, tt.expect, func(ctx context.Context, service *userService) (proto.Message, error) {
			return service.CreateTeacher(ctx, tt.req)
		}))
	}
}
