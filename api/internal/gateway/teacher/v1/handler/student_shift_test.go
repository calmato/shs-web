package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/golang/mock/gomock"
)

func TestListEnabledStudentShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 12, 0, 0, 0, 0)
	summary := &lesson.ShiftSummary{
		Id:        1,
		YearMonth: 202202,
		Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
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
	submission := &lesson.StudentSubmission{
		StudentId:      idmock,
		ShiftSummaryId: 1,
		Decided:        true,
		CreatedAt:      now.Unix(),
		UpdatedAt:      now.Unix(),
	}
	studentShifts := []*lesson.StudentShift{
		{
			StudentId:      idmock,
			ShiftId:        1,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
		{
			StudentId:      idmock,
			ShiftId:        3,
			ShiftSummaryId: 1,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		studentID string
		summaryID string
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentOut := &lesson.GetStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
			},
			studentID: idmock,
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.StudentShiftsResponse{
					Summary: &entity.StudentSubmission{
						ShiftSummaryID:   1,
						Year:             2022,
						Month:            2,
						ShiftStatus:      entity.ShiftStatusAccepting,
						SubmissionStatus: entity.StudentSubmissionStatusSubmitted,
						OpenAt:           jst.Date(2022, 1, 1, 0, 0, 0, 0),
						EndAt:            jst.Date(2022, 1, 15, 0, 0, 0, 0),
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					Shifts: entity.StudentShiftDetails{
						{
							Date:     "20220201",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 1, Enabled: true, StartTime: "1700", EndTime: "1830"},
							},
						},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 3, Enabled: true, StartTime: "1700", EndTime: "1830"},
							},
						},
					},
				},
			},
		},
		{
			name: "failed to invalid summary id",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
			},
			studentID: idmock,
			summaryID: "aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentOut := &lesson.GetStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
			},
			studentID: idmock,
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentOut := &lesson.GetStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(nil, errmock)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
			},
			studentID: idmock,
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list student shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(nil, errmock)
			},
			studentID: idmock,
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to shifts group by date",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{
					Shifts: []*lesson.Shift{{Date: "20220200"}},
				}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentOut := &lesson.GetStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
			},
			studentID: idmock,
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/shifts/%s/students/%s", tt.summaryID, tt.studentID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
