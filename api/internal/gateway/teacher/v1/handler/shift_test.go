package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestListShiftSummaries(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 12, 0, 0, 0, 0)
	summaries := []*lesson.ShiftSummary{
		{
			Id:        1,
			YearMonth: 202202,
			Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
			OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
			EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
			CreatedAt: now.Unix(),
			UpdatedAt: now.Unix(),
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
				in := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				out := &lesson.ListShiftSummariesResponse{Summaries: summaries}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), in).Return(out, nil)
			},
			query: "?status=2",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.ShiftSummariesResponse{
					Summaries: entity.ShiftSummaries{
						{
							ID:        1,
							Year:      2022,
							Month:     2,
							Status:    entity.ShiftStatusAccepting,
							OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
							EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
							CreatedAt: now,
							UpdatedAt: now,
						},
					},
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
			name:  "failed to invalid status",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?status=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list shift summaries",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), in).Return(nil, errmock)
			},
			query: "?status=2",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts%s", tt.query)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestListShiftSubmissions(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 12, 0, 0, 0, 0)
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
	// teachers := []*user.Teacher{
	// 	{
	// 		Id:            "teacherid",
	// 		LastName:      "中村",
	// 		FirstName:     "広大",
	// 		LastNameKana:  "なかむら",
	// 		FirstNameKana: "こうだい",
	// 		Mail:          "teacher-test001@calmato.jp",
	// 		Role:          user.Role_ROLE_TEACHER,
	// 		CreatedAt:     now.Unix(),
	// 		UpdatedAt:     now.Unix(),
	// 	},
	// }
	teacherSubjects := []*classroom.TeacherSubject{
		{
			TeacherId:  "teacherid",
			SubjectIds: []int64{1},
		},
	}
	teacherShifts := []*lesson.TeacherShift{
		{
			TeacherId:      "teacherid",
			ShiftId:        1,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
		{
			TeacherId:      "teacherid",
			ShiftId:        3,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	// students := []*user.Student{
	// 	{
	// 		Id:            "studentid",
	// 		LastName:      "中村",
	// 		FirstName:     "広大",
	// 		LastNameKana:  "なかむら",
	// 		FirstNameKana: "こうだい",
	// 		Mail:          "student-test001@calmato.jp",
	// 		SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
	// 		Grade:         3,
	// 		CreatedAt:     now.Unix(),
	// 		UpdatedAt:     now.Unix(),
	// 	},
	// }
	studentSubjects := []*classroom.StudentSubject{
		{
			StudentId:  "studentid",
			SubjectIds: []int64{1},
		},
	}
	studentShifts := []*lesson.StudentShift{
		{
			StudentId:      "studentid",
			ShiftId:        1,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
		{
			StudentId:      "studentid",
			ShiftId:        3,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	// lessons := []*lesson.Lesson{
	// 	{
	// 		Id:             1,
	// 		ShiftSummaryId: 1,
	// 		ShiftId:        1,
	// 		SubjectId:      1,
	// 		RoomId:         1,
	// 		TeacherId:      "teacherid",
	// 		StudentId:      "studentid",
	// 		Notes:          "感想",
	// 		CreatedAt:      now.Unix(),
	// 		UpdatedAt:      now.Unix(),
	// 	},
	// }
	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		summaryID string
		shiftID   string
		expect    *testResponse
	}{
		{
			name: "succcess",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				submissionsIn := &lesson.ListSubmissionsRequest{ShiftId: 1}
				submissionsOut := &lesson.ListSubmissionsResponse{
					TeacherShifts: teacherShifts,
					StudentShifts: studentShifts,
				}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				mocks.lesson.EXPECT().ListSubmissions(gomock.Any(), submissionsIn).Return(submissionsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
			},
			summaryID: "1",
			shiftID:   "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.ShiftSubmissionsResponse{
					Teachers: entity.Teachers{
						// {
						// 	ID:            "teacherid",
						// 	LastName:      "中村",
						// 	FirstName:     "広大",
						// 	LastNameKana:  "なかむら",
						// 	FirstNameKana: "こうだい",
						// 	Mail:          "teacher-test001@calmato.jp",
						// 	Role:          entity.RoleTeacher,
						// 	CreatedAt:     now,
						// 	UpdatedAt:     now,
						// 	Subjects: map[entity.SchoolType]entity.Subjects{
						// 		entity.SchoolTypeElementarySchool: {},
						// 		entity.SchoolTypeJuniorHighSchool: {},
						// 		entity.SchoolTypeHighSchool: {
						// 			{
						// 				ID:         1,
						// 				Name:       "国語",
						// 				Color:      "#F8BBD0",
						// 				SchoolType: entity.SchoolTypeHighSchool,
						// 				CreatedAt:  now,
						// 				UpdatedAt:  now,
						// 			},
						// 		},
						// 	},
						// },
					},
					Students: entity.Students{
						// {
						// 	ID:            "studentid",
						// 	LastName:      "中村",
						// 	FirstName:     "広大",
						// 	LastNameKana:  "なかむら",
						// 	FirstNameKana: "こうだい",
						// 	Mail:          "student-test001@calmato.jp",
						// 	SchoolType:    entity.SchoolTypeHighSchool,
						// 	Grade:         3,
						// 	CreatedAt:     now,
						// 	UpdatedAt:     now,
						// 	Subjects: entity.Subjects{
						// 		{
						// 			ID:         1,
						// 			Name:       "国語",
						// 			Color:      "#F8BBD0",
						// 			SchoolType: entity.SchoolTypeHighSchool,
						// 			CreatedAt:  now,
						// 			UpdatedAt:  now,
						// 		},
						// 	},
						// },
					},
					Lessons: entity.Lessons{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s/submissions/%s", tt.summaryID, tt.shiftID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpdateShiftSummarySchedule(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		req     *request.UpdateShiftSummaryScheduleRequest
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpdateShiftSummaryScheduleRequest{
					Id:     1,
					OpenAt: jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:  jst.Date(2022, 1, 14, 23, 59, 59, int(time.Second-time.Nanosecond)).Unix(),
				}
				out := &lesson.UpdateShiftSummaryShceduleResponse{}
				mocks.lesson.EXPECT().UpdateShiftSummarySchedule(gomock.Any(), in).Return(out, nil)
			},
			shiftID: "1",
			req: &request.UpdateShiftSummaryScheduleRequest{
				OpenDate: "20220101",
				EndDate:  "20220114",
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:    "failed to parse shift id",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "aaa",
			req: &request.UpdateShiftSummaryScheduleRequest{
				OpenDate: "20220101",
				EndDate:  "20220114",
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:    "failed to parse open at",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "1",
			req: &request.UpdateShiftSummaryScheduleRequest{
				OpenDate: "20220100",
				EndDate:  "20220114",
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:    "failed to parse end at",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "1",
			req: &request.UpdateShiftSummaryScheduleRequest{
				OpenDate: "20220101",
				EndDate:  "20220100",
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update shift summary schedule",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpdateShiftSummaryScheduleRequest{
					Id:     1,
					OpenAt: jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:  jst.Date(2022, 1, 14, 23, 59, 59, int(time.Second-time.Nanosecond)).Unix(),
				}
				mocks.lesson.EXPECT().UpdateShiftSummarySchedule(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID: "1",
			req: &request.UpdateShiftSummaryScheduleRequest{
				OpenDate: "20220101",
				EndDate:  "20220114",
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
			path := fmt.Sprintf("/v1/shifts/%s/schedule", tt.shiftID)
			req := newHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestDeleteShiftSummary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.DeleteShiftSummaryRequest{Id: 1}
				out := &lesson.DeleteShiftSummaryResponse{}
				mocks.lesson.EXPECT().DeleteShiftSummary(gomock.Any(), in).Return(out, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name:    "failed to parse shift id",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to delete shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.DeleteShiftSummaryRequest{Id: 1}
				mocks.lesson.EXPECT().DeleteShiftSummary(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s", tt.shiftID)
			req := newHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestListShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	summary := &lesson.ShiftSummary{
		Id:        1,
		YearMonth: 202202,
		Status:    lesson.ShiftStatus_SHIFT_STATUS_WAITING,
		OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
		EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
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
		{
			Id:             2,
			ShiftSummaryId: 1,
			Date:           "20220201",
			StartTime:      "1830",
			EndTime:        "2000",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
		{
			Id:             3,
			ShiftSummaryId: 1,
			Date:           "20220203",
			StartTime:      "1700",
			EndTime:        "1830",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
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
	teacherSubjects := []*classroom.TeacherSubject{
		{
			TeacherId:  "teacherid",
			SubjectIds: []int64{1},
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
	studentSubjects := []*classroom.StudentSubject{
		{
			StudentId:  "studentid",
			SubjectIds: []int64{1},
		},
	}
	studentSubmissions := []*lesson.StudentSubmission{
		{
			StudentId:        "studentid",
			ShiftSummaryId:   1,
			Decided:          true,
			SuggestedClasses: 8,
			CreatedAt:        now.Unix(),
			UpdatedAt:        now.Unix(),
		},
	}
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

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.ShiftsResponse{
					Summary: &entity.ShiftSummary{
						ID:        1,
						Year:      2022,
						Month:     2,
						Status:    entity.ShiftStatusWaiting,
						OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
						EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
						CreatedAt: now,
						UpdatedAt: now,
					},
					Shifts: entity.ShiftDetails{
						{
							Date:     "20220201",
							IsClosed: false,
							Lessons: entity.Shifts{
								{ID: 1, StartTime: "1700", EndTime: "1830"},
								{ID: 2, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.Shifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.Shifts{
								{ID: 3, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.Shifts{}},
					},
					Rooms: 4,
					Teachers: entity.TeacherSubmissionDetails{
						{
							Teacher: &entity.Teacher{
								ID:            "teacherid",
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
							LessonTotal: 1,
						},
					},
					Students: entity.StudentSubmissionDetails{
						{
							Student: &entity.Student{
								ID:            "studentid",
								LastName:      "中村",
								FirstName:     "広大",
								LastNameKana:  "なかむら",
								FirstNameKana: "こうだい",
								Mail:          "student-test001@calmato.jp",
								SchoolType:    entity.SchoolTypeHighSchool,
								Grade:         3,
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
							LessonTotal:           1,
							SuggestedClassesTotal: 8,
						},
					},
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
				},
			},
		},
		{
			name:    "failed to invalid shift id",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list teachers",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(nil, errmock)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list students",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(nil, errmock)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get rooms total",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(nil, errmock)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to mutli get teacher subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).Return(nil, errmock)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to mutli get student subjects",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(nil, errmock)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list student submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(nil, errmock)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to new shift details",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{
					Shifts: []*lesson.Shift{
						{
							Id:             1,
							ShiftSummaryId: 1,
							Date:           "20220200",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(nil, errmock)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to new lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				teachersIn := &user.ListTeachersRequest{Limit: 0, Offset: 0}
				teachersOut := &user.ListTeachersResponse{Teachers: teachers}
				studentsIn := &user.ListStudentsRequest{Limit: 0, Offset: 0}
				studentsOut := &user.ListStudentsResponse{Students: students}
				roomsIn := &classroom.GetRoomsTotalRequest{}
				roomsOut := &classroom.GetRoomsTotalResponse{Total: 4}
				teacherSubjectsIn := &classroom.MultiGetTeacherSubjectsRequest{TeacherIds: []string{"teacherid"}}
				teacherSubjectsOut := &classroom.MultiGetTeacherSubjectsResponse{
					TeacherSubjects: teacherSubjects,
					Subjects:        subjects,
				}
				studentSubjectsIn := &classroom.MultiGetStudentSubjectsRequest{StudentIds: []string{"studentid"}}
				studentSubjectsOut := &classroom.MultiGetStudentSubjectsResponse{
					StudentSubjects: studentSubjects,
					Subjects:        subjects,
				}
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: []*lesson.Shift{}}
				studentSubmissionsIn := &lesson.ListStudentSubmissionsByStudentIDsRequest{
					StudentIds: []string{"studentid"}, ShiftSummaryId: 1,
				}
				studentSubmissionsOut := &lesson.ListStudentSubmissionsByStudentIDsResponse{Submissions: studentSubmissions}
				lessonsIn := &lesson.ListLessonsByShiftSummaryIDRequest{ShiftSummaryId: 1}
				lessonsOut := &lesson.ListLessonsByShiftSummaryIDResponse{Lessons: lessons}
				mocks.user.EXPECT().ListTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
				mocks.user.EXPECT().ListStudents(gomock.Any(), studentsIn).Return(studentsOut, nil)
				mocks.classroom.EXPECT().GetRoomsTotal(gomock.Any(), roomsIn).Return(roomsOut, nil)
				mocks.classroom.EXPECT().MultiGetTeacherSubjects(gomock.Any(), teacherSubjectsIn).
					Return(teacherSubjectsOut, nil)
				mocks.classroom.EXPECT().MultiGetStudentSubjects(gomock.Any(), studentSubjectsIn).
					Return(studentSubjectsOut, nil)
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().ListStudentSubmissionsByStudentIDs(gomock.Any(), studentSubmissionsIn).
					Return(studentSubmissionsOut, nil)
				mocks.lesson.EXPECT().ListLessonsByShiftSummaryID(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
			},
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s", tt.shiftID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestCreateShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	summary := &lesson.ShiftSummary{
		Id:        1,
		YearMonth: 202202,
		Status:    lesson.ShiftStatus_SHIFT_STATUS_WAITING,
		OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
		EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
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
		{
			Id:             2,
			ShiftSummaryId: 1,
			Date:           "20220201",
			StartTime:      "1830",
			EndTime:        "2000",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
		{
			Id:             3,
			ShiftSummaryId: 1,
			Date:           "20220203",
			StartTime:      "1700",
			EndTime:        "1830",
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.CreateShiftsRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.CreateShiftsRequest{
					YearMonth:   202202,
					OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, int(time.Second-time.Nanosecond)).Unix(),
					ClosedDates: []string{"20210202", "20210214"},
				}
				out := &lesson.CreateShiftsResponse{
					Summary: summary,
					Shifts:  shifts,
				}
				mocks.lesson.EXPECT().CreateShifts(gomock.Any(), in).Return(out, nil)
			},
			req: &request.CreateShiftsRequest{
				YearMonth:   "202202",
				OpenDate:    "20220101",
				EndDate:     "20220114",
				ClosedDates: []string{"20210202", "20210214"},
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.ShiftsResponse{
					Summary: &entity.ShiftSummary{
						ID:        1,
						Year:      2022,
						Month:     2,
						Status:    entity.ShiftStatusWaiting,
						OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
						EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
						CreatedAt: now,
						UpdatedAt: now,
					},
					Shifts: entity.ShiftDetails{
						{
							Date:     "20220201",
							IsClosed: false,
							Lessons: entity.Shifts{
								{ID: 1, StartTime: "1700", EndTime: "1830"},
								{ID: 2, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.Shifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.Shifts{
								{ID: 3, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.Shifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.Shifts{}},
					},
				},
			},
		},
		{
			name:  "failed to parse year month",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateShiftsRequest{
				YearMonth:   "aaaaaa",
				OpenDate:    "20220101",
				EndDate:     "20220114",
				ClosedDates: []string{"20210202", "20210214"},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to parse open date",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateShiftsRequest{
				YearMonth:   "202202",
				OpenDate:    "20220100",
				EndDate:     "20220114",
				ClosedDates: []string{"20210202", "20210214"},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to parse end date",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			req: &request.CreateShiftsRequest{
				YearMonth:   "202202",
				OpenDate:    "20220101",
				EndDate:     "20220132",
				ClosedDates: []string{"20210202", "20210214"},
			},
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.CreateShiftsRequest{
					YearMonth:   202202,
					OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, int(time.Second-time.Nanosecond)).Unix(),
					ClosedDates: []string{"20210202", "20210214"},
				}
				mocks.lesson.EXPECT().CreateShifts(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.CreateShiftsRequest{
				YearMonth:   "202202",
				OpenDate:    "20220101",
				EndDate:     "20220114",
				ClosedDates: []string{"20210202", "20210214"},
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to new shift details",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.CreateShiftsRequest{
					YearMonth:   202202,
					OpenAt:      jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
					EndAt:       jst.Date(2022, 1, 14, 23, 59, 59, int(time.Second-time.Nanosecond)).Unix(),
					ClosedDates: []string{"20210202", "20210214"},
				}
				out := &lesson.CreateShiftsResponse{
					Summary: summary,
					Shifts: []*lesson.Shift{
						{
							Id:             1,
							ShiftSummaryId: 1,
							Date:           "20220200",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				}
				mocks.lesson.EXPECT().CreateShifts(gomock.Any(), in).Return(out, nil)
			},
			req: &request.CreateShiftsRequest{
				YearMonth:   "202202",
				OpenDate:    "20220101",
				EndDate:     "20220114",
				ClosedDates: []string{"20210202", "20210214"},
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
			path := "/v1/shifts"
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
