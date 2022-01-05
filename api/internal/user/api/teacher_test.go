package api

import (
	"context"
	"testing"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListTeachers(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &user.ListTeachersRequest{
		Limit:  30,
		Offset: 0,
	}
	params := &database.ListTeachersParams{
		Limit:  30,
		Offset: 0,
	}
	teachers := entity.Teachers{
		{
			ID:            "kSByoE6FetnPs5Byk3a9Zx",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "teacher-test001@calmato.jp",
			Role:          entity.RoleTeacher,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.ListTeachersRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeachers(req).Return(nil)
				mocks.db.Teacher.EXPECT().List(gomock.Any(), params).Return(teachers, nil)
				mocks.db.Teacher.EXPECT().Count(gomock.Any()).Return(int64(1), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.ListTeachersResponse{
					Teachers: []*user.Teacher{
						{
							Id:            "kSByoE6FetnPs5Byk3a9Zx",
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "teacher-test001@calmato.jp",
							Role:          user.Role_ROLE_TEACHER,
							CreatedAt:     now.Unix(),
							UpdatedAt:     now.Unix(),
						},
					},
					Total: 1,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.ListTeachersRequest{}
				mocks.validator.EXPECT().ListTeachers(req).Return(validation.ErrRequestValidation)
			},
			req: &user.ListTeachersRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list teachers",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeachers(req).Return(nil)
				mocks.db.Teacher.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.Teacher.EXPECT().Count(gomock.Any()).Return(int64(3), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "failed to count",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListTeachers(req).Return(nil)
				mocks.db.Teacher.EXPECT().List(gomock.Any(), params).Return(teachers, nil)
				mocks.db.Teacher.EXPECT().Count(gomock.Any()).Return(int64(0), errmock)
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
			return service.ListTeachers(ctx, tt.req)
		}))
	}
}

func TestGetTeacher(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &user.GetTeacherRequest{
		Id: "kSByoE6FetnPs5Byk3a9Zx",
	}
	teacher := &entity.Teacher{
		ID:            "kSByoE6FetnPs5Byk3a9Zx",
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          "teacher-test001@calmato.jp",
		Role:          entity.RoleTeacher,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.GetTeacherRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Get(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(teacher, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.GetTeacherResponse{
					Teacher: &user.Teacher{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          user.Role_ROLE_TEACHER,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.GetTeacherRequest{}
				mocks.validator.EXPECT().GetTeacher(req).Return(validation.ErrRequestValidation)
			},
			req: &user.GetTeacherRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().GetTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Get(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(nil, errmock)
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
			return service.GetTeacher(ctx, tt.req)
		}))
	}
}

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
			name: "failed to invalid role",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.CreateTeacherRequest{Role: user.Role(-1)}
				mocks.validator.EXPECT().CreateTeacher(req).Return(nil)
			},
			req: &user.CreateTeacherRequest{
				Role: user.Role(-1),
			},
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

func TestUpdateTeacherMail(t *testing.T) {
	t.Parallel()

	req := &user.UpdateTeacherMailRequest{
		Id:   "kSByoE6FetnPs5Byk3a9Zx",
		Mail: "teacher-test001@calmato.jp",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.UpdateTeacherMailRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherMail(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdateMail(ctx, "kSByoE6FetnPs5Byk3a9Zx", "teacher-test001@calmato.jp").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.UpdateTeacherMailResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.UpdateTeacherMailRequest{}
				mocks.validator.EXPECT().UpdateTeacherMail(req).Return(validation.ErrRequestValidation)
			},
			req: &user.UpdateTeacherMailRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update teacher mail",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherMail(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdateMail(ctx, "kSByoE6FetnPs5Byk3a9Zx", "teacher-test001@calmato.jp").Return(errmock)
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
			return service.UpdateTeacherMail(ctx, tt.req)
		}))
	}
}

func TestUpdateTeacherPassword(t *testing.T) {
	t.Parallel()

	req := &user.UpdateTeacherPasswordRequest{
		Id:                   "kSByoE6FetnPs5Byk3a9Zx",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.UpdateTeacherPasswordRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherPassword(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdatePassword(ctx, "kSByoE6FetnPs5Byk3a9Zx", "12345678").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.UpdateTeacherPasswordResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.UpdateTeacherPasswordRequest{}
				mocks.validator.EXPECT().UpdateTeacherPassword(req).Return(validation.ErrRequestValidation)
			},
			req: &user.UpdateTeacherPasswordRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update teacher password",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherPassword(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdatePassword(ctx, "kSByoE6FetnPs5Byk3a9Zx", "12345678").Return(errmock)
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
			return service.UpdateTeacherPassword(ctx, tt.req)
		}))
	}
}

func TestUpdateTeacherRole(t *testing.T) {
	t.Parallel()

	req := &user.UpdateTeacherRoleRequest{
		Id:   "kSByoE6FetnPs5Byk3a9Zx",
		Role: user.Role_ROLE_ADMINISTRATOR,
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.UpdateTeacherRoleRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherRole(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdateRole(ctx, "kSByoE6FetnPs5Byk3a9Zx", entity.RoleAdministrator).Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.UpdateTeacherRoleResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.UpdateTeacherRoleRequest{}
				mocks.validator.EXPECT().UpdateTeacherRole(req).Return(validation.ErrRequestValidation)
			},
			req: &user.UpdateTeacherRoleRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update teacher role",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateTeacherRole(req).Return(nil)
				mocks.db.Teacher.EXPECT().UpdateRole(ctx, "kSByoE6FetnPs5Byk3a9Zx", entity.RoleAdministrator).Return(errmock)
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
			return service.UpdateTeacherRole(ctx, tt.req)
		}))
	}
}

func TestDeleteTeacher(t *testing.T) {
	t.Parallel()

	req := &user.DeleteTeacherRequest{
		Id: "kSByoE6FetnPs5Byk3a9Zx",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.DeleteTeacherRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Delete(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.DeleteTeacherResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.DeleteTeacherRequest{}
				mocks.validator.EXPECT().DeleteTeacher(req).Return(validation.ErrRequestValidation)
			},
			req: &user.DeleteTeacherRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to delete teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteTeacher(req).Return(nil)
				mocks.db.Teacher.EXPECT().Delete(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(errmock)
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
			return service.DeleteTeacher(ctx, tt.req)
		}))
	}
}
