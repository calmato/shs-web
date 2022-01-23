package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestGetStudent(t *testing.T) {
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
		Grade:         1,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
	}
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
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		studentID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				subjectsOut := &classroom.GetStudentSubjectResponse{Subjects: subjects}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
			},
			studentID: idmock,
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.StudentResponse{
					Student: &entity.Student{
						ID:            idmock,
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "student-test001@calmato.jp",
						SchoolType:    3,
						Grade:         1,
						CreatedAt:     now,
						UpdatedAt:     now,
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
		},
		{
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				subjectsOut := &classroom.GetStudentSubjectResponse{Subjects: subjects}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
			},
			studentID: idmock,
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get student subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(nil, errmock)
			},
			studentID: idmock,
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/students/%s", tt.studentID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
