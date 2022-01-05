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

func TestCreateSubject(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	subject := &classroom.Subject{
		Id:         1,
		Name:       "国語",
		Color:      "#F8BBD0",
		SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.CreateSubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.CreateSubjectRequest{
					Name:       "国語",
					Color:      "#f8bbd0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				out := &classroom.CreateSubjectResponse{Subject: subject}
				mocks.classroom.EXPECT().CreateSubject(gomock.Any(), in).Return(out, nil)
			},
			req: &request.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeHighSchool,
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubjectResponse{
					Subject: &entity.Subject{
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
		{
			name:  "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeUnknown,
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.CreateSubjectRequest{
					Name:       "国語",
					Color:      "#f8bbd0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				mocks.classroom.EXPECT().CreateSubject(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeHighSchool,
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
			path := "/v1/subjects"
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateSubject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		subjectID string
		req       *request.UpdateSubjectRequest
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateSubjectRequest{
					Id:         1,
					Name:       "国語",
					Color:      "#f8bbd0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				out := &classroom.UpdateSubjectResponse{}
				mocks.classroom.EXPECT().UpdateSubject(gomock.Any(), in).Return(out, nil)
			},
			subjectID: "1",
			req: &request.UpdateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeHighSchool,
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:      "failed to invalid subject id",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			subjectID: "a",
			req: &request.UpdateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeUnknown,
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:      "failed to invalid school type",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			subjectID: "1",
			req: &request.UpdateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeUnknown,
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpdateSubjectRequest{
					Id:         1,
					Name:       "国語",
					Color:      "#f8bbd0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				mocks.classroom.EXPECT().UpdateSubject(gomock.Any(), in).Return(nil, errmock)
			},
			subjectID: "1",
			req: &request.UpdateSubjectRequest{
				Name:       "国語",
				Color:      "#f8bbd0",
				SchoolType: entity.SchoolTypeHighSchool,
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
			path := fmt.Sprintf("/v1/subjects/%s", tt.subjectID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestDeleteSubject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		subjectID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.DeleteSubjectRequest{Id: 1}
				out := &classroom.DeleteSubjectResponse{}
				mocks.classroom.EXPECT().DeleteSubject(gomock.Any(), in).Return(out, nil)
			},
			subjectID: "1",
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:      "failed to invalid subject id",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			subjectID: "a",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to delete subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.DeleteSubjectRequest{Id: 1}
				mocks.classroom.EXPECT().DeleteSubject(gomock.Any(), in).Return(nil, errmock)
			},
			subjectID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/subjects/%s", tt.subjectID)
			req := newHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
