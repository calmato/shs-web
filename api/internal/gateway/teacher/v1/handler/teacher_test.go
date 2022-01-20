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
	teacherSubjects := []*classroom.TeacherSubject{
		{
			TeacherId:  idmock,
			SubjectIds: []int64{1},
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
				teachersIn := &user.ListTeachersRequest{Limit: 100, Offset: 20}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers, Total: 1}
				subjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{idmock}}
				subjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
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
					Total: 1,
				},
			},
		},
		{
			name:  "failed to invalid limit",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?limit=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?offset=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list teachers",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 30, Offset: 0}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(nil, errmock)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 30, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers, Total: 1}
				subjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{idmock}}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), subjectsIn).Return(nil, errmock)
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
		teacherID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				subjectsOut := &classroom.GetTeacherSubjectResponse{Subjects: subjects}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
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
		},
		{
			name: "failed to get teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				subjectsOut := &classroom.GetTeacherSubjectResponse{Subjects: subjects}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(nil, errmock)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(subjectsOut, nil)
			},
			teacherID: idmock,
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get teacher subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teacherIn := &user.GetTeacherRequest{Id: idmock}
				teacherOut := &user.GetTeacherResponse{Teacher: teacher}
				subjectsIn := &classroom.GetTeacherSubjectRequest{TeacherId: idmock}
				mocks.user.EXPECT().GetTeacher(gomock.Any(), teacherIn).Return(teacherOut, nil)
				mocks.classroom.EXPECT().GetTeacherSubject(gomock.Any(), subjectsIn).Return(nil, errmock)
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
				Role:                 entity.RoleTeacher,
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
						Subjects: map[entity.SchoolType]entity.Subjects{
							entity.SchoolTypeElementarySchool: {},
							entity.SchoolTypeJuniorHighSchool: {},
							entity.SchoolTypeHighSchool:       {},
						},
					},
				},
			},
		},
		{
			name:  "failed to invalid role",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateTeacherRequest{
				Role: entity.RoleUnknown,
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
				Role:                 entity.RoleTeacher,
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

func TestUpdateTeacherMail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		req       *request.UpdateTeacherMailRequest
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateTeacherMailRequest{
					Id:   idmock,
					Mail: "teacher-test001@calmato.jp",
				}
				out := &user.UpdateTeacherMailResponse{}
				mocks.user.EXPECT().UpdateTeacherMail(gomock.Any(), in).Return(out, nil)
			},
			teacherID: idmock,
			req: &request.UpdateTeacherMailRequest{
				Mail: "teacher-test001@calmato.jp",
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
					Mail: "teacher-test001@calmato.jp",
				}
				mocks.user.EXPECT().UpdateTeacherMail(gomock.Any(), in).Return(nil, errmock)
			},
			teacherID: idmock,
			req: &request.UpdateTeacherMailRequest{
				Mail: "teacher-test001@calmato.jp",
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
			path := fmt.Sprintf("/v1/teachers/%s/mail", tt.teacherID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateTeacherPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		req       *request.UpdateTeacherPasswordRequest
		expect    *testResponse
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
			teacherID: idmock,
			req: &request.UpdateTeacherPasswordRequest{
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
			teacherID: idmock,
			req: &request.UpdateTeacherPasswordRequest{
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
			path := fmt.Sprintf("/v1/teachers/%s/password", tt.teacherID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateTeacherRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		req       *request.UpdateTeacherRoleRequest
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateTeacherRoleRequest{
					Id:   idmock,
					Role: user.Role_ROLE_ADMINISTRATOR,
				}
				out := &user.UpdateTeacherRoleResponse{}
				mocks.user.EXPECT().UpdateTeacherRole(gomock.Any(), in).Return(out, nil)
			},
			teacherID: idmock,
			req: &request.UpdateTeacherRoleRequest{
				Role: entity.RoleAdministrator,
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:      "invalid teacher role",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			teacherID: idmock,
			req: &request.UpdateTeacherRoleRequest{
				Role: entity.RoleUnknown,
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update teacher role",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.UpdateTeacherRoleRequest{
					Id:   idmock,
					Role: user.Role_ROLE_ADMINISTRATOR,
				}
				mocks.user.EXPECT().UpdateTeacherRole(gomock.Any(), in).Return(nil, errmock)
			},
			teacherID: idmock,
			req: &request.UpdateTeacherRoleRequest{
				Role: entity.RoleAdministrator,
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
			path := fmt.Sprintf("/v1/teachers/%s/role", tt.teacherID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateTeacherSubject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		req       *request.UpdateTeacherSubjectRequest
		expect    *testResponse
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
			teacherID: idmock,
			req: &request.UpdateTeacherSubjectRequest{
				SchoolType: entity.SchoolTypeHighSchool,
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:      "failed to invalid school type",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			teacherID: idmock,
			req: &request.UpdateTeacherSubjectRequest{
				SchoolType: entity.SchoolTypeUnknown,
				SubjectIDs: []int64{1},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to upsert teacher subject",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &classroom.UpsertTeacherSubjectRequest{
					TeacherId:  idmock,
					SubjectIds: []int64{1},
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				}
				mocks.classroom.EXPECT().UpsertTeacherSubject(gomock.Any(), in).Return(nil, errmock)
			},
			teacherID: idmock,
			req: &request.UpdateTeacherSubjectRequest{
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
			path := fmt.Sprintf("/v1/teachers/%s/subjects", tt.teacherID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestDeleteTeacher(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.DeleteTeacherRequest{Id: idmock}
				out := &user.DeleteTeacherResponse{}
				mocks.user.EXPECT().DeleteTeacher(gomock.Any(), in).Return(out, nil)
			},
			teacherID: idmock,
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to delete teacher",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &user.DeleteTeacherRequest{Id: idmock}
				mocks.user.EXPECT().DeleteTeacher(gomock.Any(), in).Return(nil, errmock)
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
			req := newHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
