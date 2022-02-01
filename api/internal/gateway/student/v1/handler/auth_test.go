package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
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
	studentSubject := &classroom.StudentSubject{
		StudentId:  idmock,
		SubjectIds: []int64{1},
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
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				subjectsOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject, Subjects: subjects}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
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
			name: "failed to get student",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				subjectsOut := &classroom.GetStudentSubjectResponse{StudentSubject: studentSubject, Subjects: subjects}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get student subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				studentIn := &user.GetStudentRequest{Id: idmock}
				studentOut := &user.GetStudentResponse{Student: student}
				subjectsIn := &classroom.GetStudentSubjectRequest{StudentId: idmock}
				mocks.user.EXPECT().GetStudent(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.classroom.EXPECT().GetStudentSubject(gomock.Any(), subjectsIn).Return(nil, errmock)
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

func TestUpdateMyMail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateMyMailRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateStudentMailRequest{
					Id:   idmock,
					Mail: "student-test01@calmato.jp",
				}
				out := &user.UpdateStudentMailResponse{}
				mocks.user.EXPECT().UpdateStudentMail(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMyMailRequest{
				Mail: "student-test01@calmato.jp",
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to update student mail",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateStudentMailRequest{
					Id:   idmock,
					Mail: "student-test01@calmato.jp",
				}
				mocks.user.EXPECT().UpdateStudentMail(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateMyMailRequest{
				Mail: "student-test01@calmato.jp",
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
			const path = "/v1/me/mail"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateMyPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateMyPasswordRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateStudentPasswordRequest{
					Id:                   idmock,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				out := &user.UpdateStudentPasswordResponse{}
				mocks.user.EXPECT().UpdateStudentPassword(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMyPasswordRequest{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to update student password",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateStudentPasswordRequest{
					Id:                   idmock,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				mocks.user.EXPECT().UpdateStudentPassword(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateMyPasswordRequest{
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
			const path = "/v1/me/password"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
