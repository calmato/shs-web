package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/golang/mock/gomock"
)

func TestGetAuth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		userID string
		expect *testResponse
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:            idMock,
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          entity.Role(entity.RoleTeacher),
					},
				},
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
