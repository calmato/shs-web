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
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestListLessons(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 15, 0, 0, 0, 0)
	lessons := []*lesson.Lesson{
		{
			Id:             1,
			ShiftSummaryId: 1,
			ShiftId:        1,
			SubjectId:      1,
			RoomId:         1,
			TeacherId:      "teacherid",
			StudentId:      "studentid",
			Notes:          "感想",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	shifts := []*lesson.Shift{
		{
			Id:             1,
			ShiftSummaryId: 1,
			Date:           "20220201",
			StartTime:      "1700",
			EndTime:        "1830",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	teachers := []*user.Teacher{
		{
			Id:            "teacherid",
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
	students := []*user.Student{
		{
			Id:            "studentid",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "student-test001@calmato.jp",
			SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			Grade:         3,
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
				lessonsIn := &lesson.ListLessonsByDurationRequest{
					TeacherId: idmock,
					Since:     "20220101",
					Until:     "20220131",
				}
				lessonsOut := &lesson.ListLessonsByDurationResponse{
					Lessons: lessons,
					Shifts:  shifts,
				}
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.LessonsResponse{
					Lessons: entity.Lessons{
						{
							ID:        1,
							ShiftID:   1,
							SubjectID: 1,
							Room:      1,
							TeacherID: "teacherid",
							StudentID: "studentid",
							Notes:     "感想",
							StartAt:   jst.Date(2022, 2, 1, 17, 0, 0, 0),
							EndAt:     jst.Date(2022, 2, 1, 18, 30, 0, 0),
							CreatedAt: now,
							UpdatedAt: now,
						},
					},
					Teachers: entity.Teachers{
						{
							ID:            "teacherid",
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "teacher-test001@calmato.jp",
							Role:          entity.RoleTeacher,
							Subjects: map[entity.SchoolType]entity.Subjects{
								entity.SchoolTypeElementarySchool: {},
								entity.SchoolTypeJuniorHighSchool: {},
								entity.SchoolTypeHighSchool:       {},
							},
							CreatedAt: now,
							UpdatedAt: now,
						},
					},
					Students: entity.Students{
						{
							ID:            "studentid",
							LastName:      "中村",
							FirstName:     "広大",
							LastNameKana:  "なかむら",
							FirstNameKana: "こうだい",
							Mail:          "student-test001@calmato.jp",
							SchoolType:    entity.SchoolTypeHighSchool,
							Grade:         3,
							Subjects:      entity.Subjects{},
							CreatedAt:     now,
							UpdatedAt:     now,
						},
					},
				},
			},
		},
		{
			name: "failed to list lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				lessonsIn := &lesson.ListLessonsByDurationRequest{
					TeacherId: idmock,
					Since:     "20220101",
					Until:     "20220107",
				}
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(nil, errmock)
			},
			query: "?since=20220101&until=20220107",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get teachers",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				lessonsIn := &lesson.ListLessonsByDurationRequest{
					TeacherId: idmock,
					Since:     "20220101",
					Until:     "20220131",
				}
				lessonsOut := &lesson.ListLessonsByDurationResponse{
					Lessons: lessons,
					Shifts:  shifts,
				}
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(nil, errmock)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get students",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				lessonsIn := &lesson.ListLessonsByDurationRequest{
					TeacherId: idmock,
					Since:     "20220101",
					Until:     "20220131",
				}
				lessonsOut := &lesson.ListLessonsByDurationResponse{
					Lessons: lessons,
					Shifts:  shifts,
				}
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(nil, errmock)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to new lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				lessonsIn := &lesson.ListLessonsByDurationRequest{
					TeacherId: idmock,
					Since:     "20220101",
					Until:     "20220131",
				}
				lessonsOut := &lesson.ListLessonsByDurationResponse{
					Lessons: lessons,
					Shifts:  []*lesson.Shift{},
				}
				teachersIn := &user.MultiGetTeachersRequest{Ids: []string{"teacherid"}}
				teachersOut := &user.MultiGetTeachersResponse{Teachers: teachers}
				studentsIn := &user.MultiGetStudentsRequest{Ids: []string{"studentid"}}
				studentsOut := &user.MultiGetStudentsResponse{Students: students}
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().MultiGetStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
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
			path := fmt.Sprintf("/v1/lessons%s", tt.query)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req, withNow(now))
		})
	}
}

func TestCreateLesson(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 15, 0, 0, 0, 0)
	l := &lesson.Lesson{
		Id:             1,
		ShiftSummaryId: 1,
		ShiftId:        1,
		SubjectId:      1,
		RoomId:         1,
		TeacherId:      "teacherid",
		StudentId:      "studentid",
		Notes:          "感想",
		CreatedAt:      now.Unix(),
		UpdatedAt:      now.Unix(),
	}
	shift := &lesson.Shift{
		Id:             1,
		ShiftSummaryId: 1,
		Date:           "20220201",
		StartTime:      "1700",
		EndTime:        "1830",
		CreatedAt:      now.Unix(),
		UpdatedAt:      now.Unix(),
	}

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		req     *request.CreateLessonRequest
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.CreateLessonRequest{
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				out := &lesson.CreateLessonResponse{Lesson: l, Shift: shift}
				mocks.lesson.EXPECT().CreateLesson(gomock.Any(), in).Return(out, nil)
			},
			shiftID: "1",
			req: &request.CreateLessonRequest{
				ShiftID:   1,
				SubjectID: 1,
				Room:      1,
				TeacherID: "teacherid",
				StudentID: "studentid",
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.LessonResponse{
					Lesson: &entity.Lesson{
						ID:        1,
						ShiftID:   1,
						SubjectID: 1,
						Room:      1,
						TeacherID: "teacherid",
						StudentID: "studentid",
						Notes:     "感想",
						StartAt:   jst.Date(2022, 2, 1, 17, 0, 0, 0),
						EndAt:     jst.Date(2022, 2, 1, 18, 30, 0, 0),
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
		},
		{
			name:    "failed to invalid shift summary id",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "a",
			req:     &request.CreateLessonRequest{},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create lesson",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.CreateLessonRequest{
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				mocks.lesson.EXPECT().CreateLesson(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID: "1",
			req: &request.CreateLessonRequest{
				ShiftID:   1,
				SubjectID: 1,
				Room:      1,
				TeacherID: "teacherid",
				StudentID: "studentid",
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
			path := fmt.Sprintf("/v1/shifts/%s/lessons", tt.shiftID)
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateLesson(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID  string
		lessonID string
		req      *request.UpdateLessonRequest
		expect   *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpdateLessonRequest{
					LessonId:       1,
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				out := &lesson.UpdateLessonResponse{}
				mocks.lesson.EXPECT().UpdateLesson(gomock.Any(), in).Return(out, nil)
			},
			shiftID:  "1",
			lessonID: "1",
			req: &request.UpdateLessonRequest{
				ShiftID:   1,
				SubjectID: 1,
				Room:      1,
				TeacherID: "teacherid",
				StudentID: "studentid",
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:     "failed to invalid shift summary id",
			setup:    func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID:  "a",
			lessonID: "1",
			req:      &request.UpdateLessonRequest{},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:     "failed to invalid lesson id",
			setup:    func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID:  "1",
			lessonID: "a",
			req:      &request.UpdateLessonRequest{},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update lesson",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpdateLessonRequest{
					LessonId:       1,
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				mocks.lesson.EXPECT().UpdateLesson(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID:  "1",
			lessonID: "1",
			req: &request.UpdateLessonRequest{
				ShiftID:   1,
				SubjectID: 1,
				Room:      1,
				TeacherID: "teacherid",
				StudentID: "studentid",
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
			path := fmt.Sprintf("/v1/shifts/%s/lessons/%s", tt.shiftID, tt.lessonID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestDeleteLesson(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID  string
		lessonID string
		expect   *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.DeleteLessonRequest{LessonId: 1}
				out := &lesson.DeleteLessonResponse{}
				mocks.lesson.EXPECT().DeleteLesson(gomock.Any(), in).Return(out, nil)
			},
			shiftID:  "1",
			lessonID: "1",
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to invalid lesson id",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
			},
			shiftID:  "1",
			lessonID: "a",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to delete lesson",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.DeleteLessonRequest{LessonId: 1}
				mocks.lesson.EXPECT().DeleteLesson(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID:  "1",
			lessonID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s/lessons/%s", tt.shiftID, tt.lessonID)
			req := newHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
