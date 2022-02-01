package api

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/internal/user/validation"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func TestListStudents(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &user.ListStudentsRequest{
		Limit:  30,
		Offset: 0,
	}

	params := &database.ListStudentsParams{
		Limit:  30,
		Offset: 0,
	}
	students := entity.Students{
		{
			ID:            "kSByoE6FetnPs5Byk3a9Zx",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "student-test001@calmato.jp",
			BirthYear:     2005,
			CreatedAt:     now,
			UpdatedAt:     now,
			SchoolType:    entity.SchoolTypeHighSchool,
			Grade:         1,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.ListStudentsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudents(req).Return(nil)
				mocks.db.Student.EXPECT().List(gomock.Any(), params).Return(students, nil)
				mocks.db.Student.EXPECT().Count(gomock.Any()).Return(int64(1), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.ListStudentsResponse{
					Students: []*user.Student{
						{
							Id:            "kSByoE6FetnPs5Byk3a9Zx",
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "student-test001@calmato.jp",
							CreatedAt:     now.Unix(),
							UpdatedAt:     now.Unix(),
							SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							Grade:         1,
						},
					},
					Total: 1,
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.ListStudentsRequest{}
				mocks.validator.EXPECT().ListStudents(req).Return(validation.ErrRequestValidation)
			},
			req: &user.ListStudentsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to list students",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudents(req).Return(nil)
				mocks.db.Student.EXPECT().List(gomock.Any(), params).Return(nil, errmock)
				mocks.db.Student.EXPECT().Count(gomock.Any()).Return(int64(3), nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.Internal,
			},
		},
		{
			name: "faild to count",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().ListStudents(req).Return(nil)
				mocks.db.Student.EXPECT().List(gomock.Any(), params).Return(students, nil)
				mocks.db.Student.EXPECT().Count(gomock.Any()).Return(int64(0), errmock)
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
			return service.ListStudents(ctx, tt.req)
		}))
	}
}

func TestMultiGetStudents(t *testing.T) {
	t.Parallel()
	now := jst.Now()

	req := &user.MultiGetStudentsRequest{
		Ids: []string{"kSByoE6FetnPs5Byk3a9Zx"},
	}
	students := entity.Students{
		{
			ID:            "kSByoE6FetnPs5Byk3a9Zx",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "student-test001@calmato.jp",
			SchoolType:    entity.SchoolTypeHighSchool,
			Grade:         1,
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.MultiGetStudentsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().MultiGetStudents(req).Return(nil)
				mocks.db.Student.EXPECT().MultiGet(gomock.Any(), []string{"kSByoE6FetnPs5Byk3a9Zx"}).Return(students, nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.MultiGetStudentsResponse{
					Students: []*user.Student{
						{
							Id:            "kSByoE6FetnPs5Byk3a9Zx",
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "student-test001@calmato.jp",
							CreatedAt:     now.Unix(),
							UpdatedAt:     now.Unix(),
							SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							Grade:         1,
						},
					},
				},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.MultiGetStudentsRequest{}
				mocks.validator.EXPECT().MultiGetStudents(req).Return(validation.ErrRequestValidation)
			},
			req: &user.MultiGetStudentsRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().MultiGetStudents(req).Return(nil)
				mocks.db.Student.EXPECT().MultiGet(gomock.Any(), []string{"kSByoE6FetnPs5Byk3a9Zx"}).Return(nil, errmock)
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
			return service.MultiGetStudents(ctx, tt.req)
		}))
	}
}

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
		SchoolType:    entity.SchoolTypeHighSchool,
		Grade:         1,
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
						SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						Grade:         1,
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
			name: "failed to create Student",
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

func TestUpdateStudentMail(t *testing.T) {
	t.Parallel()

	req := &user.UpdateStudentMailRequest{
		Id:   "kSByoE6FetnPs5Byk3a9Zx",
		Mail: "student-test001@calmato.jp",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.UpdateStudentMailRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateStudentMail(req).Return(nil)
				mocks.db.Student.EXPECT().UpdateMail(ctx, "kSByoE6FetnPs5Byk3a9Zx", "student-test001@calmato.jp").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.UpdateStudentMailResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.UpdateStudentMailRequest{}
				mocks.validator.EXPECT().UpdateStudentMail(req).Return(validation.ErrRequestValidation)
			},
			req: &user.UpdateStudentMailRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update student mail",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateStudentMail(req).Return(nil)
				mocks.db.Student.EXPECT().UpdateMail(ctx, "kSByoE6FetnPs5Byk3a9Zx", "student-test001@calmato.jp").Return(errmock)
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
			return service.UpdateStudentMail(ctx, tt.req)
		}))
	}
}

func TestUpdateStudentPassword(t *testing.T) {
	t.Parallel()

	req := &user.UpdateStudentPasswordRequest{
		Id:                   "kSByoE6FetnPs5Byk3a9Zx",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.UpdateStudentPasswordRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateStudentPassword(req).Return(nil)
				mocks.db.Student.EXPECT().UpdatePassword(ctx, "kSByoE6FetnPs5Byk3a9Zx", "12345678").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.UpdateStudentPasswordResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.UpdateStudentPasswordRequest{}
				mocks.validator.EXPECT().UpdateStudentPassword(req).Return(validation.ErrRequestValidation)
			},
			req: &user.UpdateStudentPasswordRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to update student password",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().UpdateStudentPassword(req).Return(nil)
				mocks.db.Student.EXPECT().UpdatePassword(ctx, "kSByoE6FetnPs5Byk3a9Zx", "12345678").Return(errmock)
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
			return service.UpdateStudentPassword(ctx, tt.req)
		}))
	}
}

func TestDeleteStudent(t *testing.T) {
	t.Parallel()

	req := &user.DeleteStudentRequest{
		Id: "kSByoE6FetnPs5Byk3a9Zx",
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks)
		req    *user.DeleteStudentRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Delete(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(nil)
			},
			req: req,
			expect: &testResponse{
				code: codes.OK,
				body: &user.DeleteStudentResponse{},
			},
		},
		{
			name: "invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				req := &user.DeleteStudentRequest{}
				mocks.validator.EXPECT().DeleteStudent(req).Return(validation.ErrRequestValidation)
			},
			req: &user.DeleteStudentRequest{},
			expect: &testResponse{
				code: codes.InvalidArgument,
			},
		},
		{
			name: "failed to delete Student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks) {
				mocks.validator.EXPECT().DeleteStudent(req).Return(nil)
				mocks.db.Student.EXPECT().Delete(ctx, "kSByoE6FetnPs5Byk3a9Zx").Return(errmock)
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
			return service.DeleteStudent(ctx, tt.req)
		}))
	}
}
