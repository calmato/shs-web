package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestGetStudent(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &user.GetStudentRequest{
		Id: "lMByuO6FetnPs5Byk3s2Jy",
	}
	student := &entity.Student{
		ID:            "lMByuO6FetnPs5Byk3s2Jy",
		LastName:      "浜田",
		FirstName:     "直志",
		LastNameKana:  "はまだ",
		FirstNameKana: "ただし",
		Mail:          "student-test001@calmato.jp",
		BirthYear:     2021,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.GetStudentRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Get(ctx, "lMByuO6FetnPs5Byk3s2Jy").Return(student, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.GetStudentResponse{
					Student: &user.Student{
						Id:            "lMByuO6FetnPs5Byk3s2Jy",
						LastName:      "浜田",
						FirstName:     "直志",
						LastNameKana:  "はまだ",
						FirstNameKana: "ただし",
						Mail:          "student-test001@calmato.jp",
						BirthYear:     2021,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.GetStudentRequest{}
				mocks.validator.EXPECT().GetStudent(req).Return(validation.ErrRequestValidation)
			},
			req: &user.GetStudentRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Get(ctx, "lMByuO6FetnPs5Byk3s2Jy").Return(nil, errmock)
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
			return service.GetStudent(ctx, tt.req)
		}))
	}
}

func TestCreateStudent(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	req := &user.CreateStudentRequest{
		LastName:      "山田",
		FirstName:     "太郎",
		LastNameKana:  "やまだ",
		FirstNameKana: "たろう",
		Mail:          "student-test001@calamto.jp",
		Password:      "12345678",
		SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
		Grade:         1,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.CreateStudentRequest
		now    time.Time
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			req: req,
			now: now,
			expect: &testResponse{
				code: codes.OK,
			},
		},
		{
			name: "invalid arugment",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.CreateStudentRequest{}
				mocks.validator.EXPECT().CreateStudent(req).Return(validation.ErrRequestValidation)
			},
			req: &user.CreateStudentRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to create teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().CreateStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Create(ctx, gomock.Any()).Return(errmock)
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
			return service.CreateStudent(ctx, tt.req)
		}))
	}
}
