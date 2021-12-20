package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestGetAuth(t *testing.T) {
	t.Parallel()
	now := jst.Now()
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
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.GetTeacherRequest{Id: idmock}
				out := &user.GetTeacherResponse{Teacher: teacher}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), in).Return(out, nil)
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:            idmock,
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          entity.RoleTeacher,
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
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			const path = "/v1/me"
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateMySubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	subjects := []*classroom.Subject{
		{
			Id:         1,
			Name:       "国語",
			Color:      "#f8bbd0",
			SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateMySubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.MultiGetSubjectsRequest{Ids: []int64{1}}
				out := &classroom.MultiGetSubjectsResponse{Subjects: subjects}
				mocks.classroom.EXPECT().MultiGetSubjects(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMySubjectRequest{
				SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.AuthResponse{},
			},
		},
		{
			name:  "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.UpdateMySubjectRequest{
				SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_UNKNOWN),
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.MultiGetSubjectsRequest{Ids: []int64{1}}
				mocks.classroom.EXPECT().MultiGetSubjects(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateMySubjectRequest{
				SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to unmatch subjects length",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.MultiGetSubjectsRequest{Ids: []int64{1}}
				out := &classroom.MultiGetSubjectsResponse{Subjects: []*classroom.Subject{}}
				mocks.classroom.EXPECT().MultiGetSubjects(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMySubjectRequest{
				SchoolType: int32(classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL),
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			const path = "/v1/me/subjects"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
