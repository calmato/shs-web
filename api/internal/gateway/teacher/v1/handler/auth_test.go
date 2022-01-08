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
	teacherSubject := &classroom.TeacherSubject{
		TeacherId:  idmock,
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
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				subjectsOut := &classroom.GetTeacherSubjectResponse{TeacherSubject: teacherSubject, Subjects: subjects}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
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
					Subjects: map[entity.SchoolType]entity.Subjects{
						entity.SchoolTypeElementarySchool: {},
						entity.SchoolTypeJuniorHighSchool: {},
						entity.SchoolTypeHighSchool: {
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
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				subjectsOut := &classroom.GetTeacherSubjectResponse{TeacherSubject: teacherSubject, Subjects: subjects}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get teacher subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(nil, errmock)
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
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpdateMySubjectRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpsertTeacherSubjectRequest{
					TeacherId:  idmock,
					SubjectIds: []int64{1},
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				out := &classroom.UpsertTeacherSubjectResponse{}
				mocks.classroom.EXPECT().UpsertTeacherSubject(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMySubjectRequest{
				SchoolType: entity.SchoolTypeHighSchool,
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:  "failed to invalid school type",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.UpdateMySubjectRequest{
				SchoolType: entity.SchoolTypeUnknown,
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update teacher subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpsertTeacherSubjectRequest{
					TeacherId:  idmock,
					SubjectIds: []int64{1},
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				mocks.classroom.EXPECT().UpsertTeacherSubject(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateMySubjectRequest{
				SchoolType: entity.SchoolTypeHighSchool,
				SubjectIDs: []int64{1},
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
			const path = "/v1/me/subjects"
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
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
				in := &user.UpdateTeacherMailRequest{
					Id:   idmock,
					Mail: "teacher-test01@calmato.jp",
				}
				out := &user.UpdateTeacherMailResponse{}
				mocks.user.EXPECT().UpdateTeacherMail(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpdateMyMailRequest{
				Mail: "teacher-test01@calmato.jp",
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to update teacher mail",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateTeacherMailRequest{
					Id:   idmock,
					Mail: "teacher-test01@calmato.jp",
				}
				mocks.user.EXPECT().UpdateTeacherMail(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpdateMyMailRequest{
				Mail: "teacher-test01@calmato.jp",
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
				in := &user.UpdateTeacherPasswordRequest{
					Id:                   idmock,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				out := &user.UpdateTeacherPasswordResponse{}
				mocks.user.EXPECT().UpdateTeacherPassword(gomock.Any(), in).Return(out, nil)
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
			name: "failed to update teacher password",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateTeacherPasswordRequest{
					Id:                   idmock,
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}
				mocks.user.EXPECT().UpdateTeacherPassword(gomock.Any(), in).Return(nil, errmock)
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
