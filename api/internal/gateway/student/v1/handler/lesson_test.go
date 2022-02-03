package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
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
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
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
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(nil, errmock)
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
				mocks.lesson.EXPECT().ListLessonsByDuration(gomock.Any(), lessonsIn).Return(lessonsOut, nil)
				mocks.user.EXPECT().MultiGetTeachers(gomock.Any(), teachersIn).Return(teachersOut, nil)
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
