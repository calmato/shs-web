package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/golang/mock/gomock"
)

func TestListSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	subjects := []*classroom.Subject{
		{
			Id:         1,
			Name:       "国語",
			Color:      "#F8BBD0",
			SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
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
				in := &classroom.ListSubjectsRequest{SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL}
				out := &classroom.ListSubjectsResponse{Subjects: subjects}
				mocks.classroom.EXPECT().ListSubjects(gomock.Any(), in).Return(out, nil)
			},
			query: "?type=3",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubjectsResponse{
					Subjects: entity.Subjects{
						{
							ID:         1,
							Name:       "国語",
							Color:      "#F8BBD0",
							SchoolType: entity.SchoolTypeHighSchool,
							CreatedAt:  now,
							UpdatedAt:  now,
						},
					},
				},
			},
		},
		{
			name:  "failed to invlalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?type=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.ListSubjectsRequest{SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL}
				mocks.classroom.EXPECT().ListSubjects(gomock.Any(), in).Return(nil, errmock)
			},
			query: "?type=3",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/subjects" + tt.query
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
