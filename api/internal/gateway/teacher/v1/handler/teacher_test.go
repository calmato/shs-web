package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestListTeachers(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	teachers := []*user.Teacher{
		{
			Id:            idmock,
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "teacher-test001@calmato.jp",
			Role:          user.Role_ROLE_TEACHER,
			CreatedAt:     now.Unix(),
			UpdatedAt:     now.Unix(),
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		query  string
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.ListTeachersRequest{Limit: 100, Offset: 20}
				out := &user.ListTeachersResponse{Teachers: teachers, Total: 1}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), in).Return(out, nil)
			},
			query: "?limit=100&offset=20",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.TeachersResponse{
					Teachers: entity.Teachers{
						{
							ID:            idmock,
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "teacher-test001@calmato.jp",
							Role:          entity.RoleTeacher,
							CreatedAt:     now,
							UpdatedAt:     now,
						},
					},
					Total: 1,
				},
			},
		},
		{
			name: "failed to list teachers",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.ListTeachersRequest{Limit: 30, Offset: 0}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), in).Return(nil, errmock)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/teachers" + tt.query
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestGetTeacher(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	teacher := &user.Teacher{
		Id:            idmock,
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          "teacher-test001@calmato.jp",
		Role:          user.Role_ROLE_TEACHER,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
	}
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.GetTeacherRequest{Id: idmock}
				out := &user.GetTeacherResponse{Teacher: teacher}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
			},
			teacherID: idmock,
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.TeacherResponse{
					Teacher: &entity.Teacher{
						ID:            idmock,
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          entity.RoleTeacher,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
				},
			},
		},
		{
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.GetTeacherRequest{Id: idmock}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(nil, errmock)
			},
			teacherID: idmock,
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/teachers/%s", tt.teacherID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestCreateTeacher(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	teacher := &user.Teacher{
		Id:            idmock,
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          "teacher-test001@calmato.jp",
		Role:          user.Role_ROLE_TEACHER,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.CreateTeacherRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.CreateTeacherRequest{
					LastName:             "中村",
					FirstName:            "広大",
					LastNameKana:         "なかむら",
					FirstNameKana:        "こうだい",
					Mail:                 "teacher-test001@calmato.jp",
					Role:                 user.Role_ROLE_TEACHER,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				out := &user.CreateTeacherResponse{Teacher: teacher}
				mocks.user.EXPECT().CreateTeacher(gomock.Any(), in).Return(out, nil)
			},
			req: &request.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 int32(entity.RoleTeacher),
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.TeacherResponse{
					Teacher: &entity.Teacher{
						ID:            idmock,
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          entity.RoleTeacher,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
				},
			},
		},
		{
			name:  "failed to invalid role",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateTeacherRequest{
				Role: int32(entity.RoleUnknown),
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.CreateTeacherRequest{
					LastName:             "中村",
					FirstName:            "広大",
					LastNameKana:         "なかむら",
					FirstNameKana:        "こうだい",
					Mail:                 "teacher-test001@calmato.jp",
					Role:                 user.Role_ROLE_TEACHER,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				mocks.user.EXPECT().CreateTeacher(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 int32(entity.RoleTeacher),
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			const path = "/v1/teachers"
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req, withNow(now))
		})
	}
}
