package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestGetAuth(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	student := &user.Student{
		Id:            idmock,
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          "student-test001@calmato.jp",
		SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
		Grade:         3,
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
				in := &user.GetStudentRequest{Id: idmock}
				out := &user.GetStudentResponse{Student: student}
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(out, nil)
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
						Mail:          "student-test001@calmato.jp",
						SchoolType:    entity.SchoolTypeHighSchool,
						Grade:         3,
					},
				},
			},
		},
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.GetStudentRequest{Id: idmock}
				mocks.user.EXPECT().GetStudent(gomock.Any(), in).Return(nil, errmock)
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
