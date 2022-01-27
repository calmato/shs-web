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

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		query   string
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.ListLessonsRequest{
					ShiftSummaryId: 1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				out := &lesson.ListLessonsResponse{
					Lessons: lessons,
					Shifts:  shifts,
					Total:   1,
				}
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), in).Return(out, nil)
			},
			shiftID: "1",
			query:   "?teacherId=teacherid&studentId=studentid",
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
					Total: 1,
				},
			},
		},
		{
			name:    "failed to invalid shift summary id",
			setup:   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			shiftID: "a",
			query:   "",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list lessons",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.ListLessonsRequest{
					ShiftSummaryId: 1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
				}
				mocks.lesson.EXPECT().ListLessons(gomock.Any(), in).Return(nil, errmock)
			},
			shiftID: "1",
			query:   "?teacherId=teacherid&studentId=studentid",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s/lessons%s", tt.shiftID, tt.query)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
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
