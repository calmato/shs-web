package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListSubmissions(t *testing.T) {
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
		{
			Id:        2,
			YearMonth: 202203,
			Status:    lesson.ShiftStatus_SHIFT_STATUS_WAITING,
			OpenAt:    jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
			EndAt:     jst.Date(2022, 2, 15, 0, 0, 0, 0).Unix(),
			CreatedAt: now.Unix(),
			UpdatedAt: now.Unix(),
		},
	}
	submissions := []*lesson.StudentSubmission{
		{
			StudentId:      idmock,
			ShiftSummaryId: 1,
			Decided:        true,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
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
				shiftsIn := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				shiftsOut := &lesson.ListShiftSummariesResponse{Summaries: summaries}
				submissionsIn := &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
					StudentId:       idmock,
					ShiftSummaryIds: []int64{1, 2},
				}
				submissionsOut := &lesson.ListStudentSubmissionsByShiftSummaryIDsResponse{
					Submissions: submissions,
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().
					ListStudentSubmissionsByShiftSummaryIDs(gomock.Any(), submissionsIn).
					Return(submissionsOut, nil)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionsResponse{
					Summaries: entity.StudentSubmissions{
						{
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
						{
							ShiftSummaryID:   2,
							Year:             2022,
							Month:            3,
							ShiftStatus:      entity.ShiftStatusWaiting,
							SubmissionStatus: entity.StudentSubmissionStatusWaiting,
							OpenAt:           jst.Date(2022, 2, 1, 0, 0, 0, 0),
							EndAt:            jst.Date(2022, 2, 15, 0, 0, 0, 0),
							CreatedAt:        now,
							UpdatedAt:        now,
						},
					},
				},
			},
		},
		{
			name:  "failed to parse limit",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?limit=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to parse offset",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			query: "?offset=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list shift summaries",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				shiftsIn := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), shiftsIn).Return(nil, errmock)
			},
			query: "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list student submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				shiftsIn := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				shiftsOut := &lesson.ListShiftSummariesResponse{Summaries: summaries}
				submissionsIn := &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
					StudentId:       idmock,
					ShiftSummaryIds: []int64{1, 2},
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().
					ListStudentSubmissionsByShiftSummaryIDs(gomock.Any(), submissionsIn).
					Return(nil, errmock)
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
			path := fmt.Sprintf("/v1/submissions%s", tt.query)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestGetSubmission(t *testing.T) {
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
		SuggestedLessons: []*lesson.SuggestedLesson{
			{SubjectId: 1, Total: 4},
		},
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
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
	template := &lesson.StudentShiftTemplate{
		StudentId: "studentid",
		Schedules: []*lesson.ShiftSchedule{
			{
				Weekday: int32(time.Thursday),
				Lessons: []*lesson.LessonSchedule{
					{StartTime: "1700", EndTime: "1830"},
				},
			},
		},
		SuggestedLessons: []*lesson.SuggestedLesson{
			{SubjectId: 1, Total: 4},
		},
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		summaryID string
		expect    *testResponse
	}{
		{
			name: "success to already submitted",
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
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionResponse{
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
								{ID: 2, Enabled: false, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.StudentShifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 3, Enabled: true, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.StudentShifts{}},
					},
					SuggestedLessons: entity.StudentSuggestedLessons{
						{SubjectID: 1, Total: 4},
					},
				},
			},
		},
		{
			name: "success to not submitted",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentErr := status.Error(codes.NotFound, errmock.Error())
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(nil, studentErr)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionResponse{
					Summary: &entity.StudentSubmission{
						ShiftSummaryID:   1,
						Year:             2022,
						Month:            2,
						ShiftStatus:      entity.ShiftStatusAccepting,
						SubmissionStatus: entity.StudentSubmissionStatusWaiting,
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
								{ID: 1, Enabled: false, StartTime: "1700", EndTime: "1830"},
								{ID: 2, Enabled: false, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.StudentShifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 3, Enabled: true, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.StudentShifts{}},
					},
					SuggestedLessons: entity.StudentSuggestedLessons{
						{SubjectID: 1, Total: 4},
					},
				},
			},
		},
		{
			name: "success to not submitted and template is empty",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.GetStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
				}
				studentErr := status.Error(codes.NotFound, errmock.Error())
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateErr := status.Error(codes.NotFound, errmock.Error())
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(nil, studentErr)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(nil, templateErr)
			},
			summaryID: "1",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionResponse{
					Summary: &entity.StudentSubmission{
						ShiftSummaryID:   1,
						Year:             2022,
						Month:            2,
						ShiftStatus:      entity.ShiftStatusAccepting,
						SubmissionStatus: entity.StudentSubmissionStatusWaiting,
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
								{ID: 1, Enabled: false, StartTime: "1700", EndTime: "1830"},
								{ID: 2, Enabled: false, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.StudentShifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 3, Enabled: false, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.StudentShifts{}},
					},
					SuggestedLessons: entity.StudentSuggestedLessons{},
				},
			},
		},
		{
			name: "failed to invalid summary id",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
			},
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
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
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
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(nil, errmock)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
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
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(nil, errmock)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
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
				studentOut := &lesson.GetStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(nil, errmock)
			},
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
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().GetStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
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
			path := fmt.Sprintf("/v1/submissions/%s", tt.summaryID)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpsertSubmission(t *testing.T) {
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
		SuggestedLessons: []*lesson.SuggestedLesson{
			{SubjectId: 1, Total: 4},
		},
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
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
		summaryID string
		req       *request.UpsertSubmissionRequest
		expect    *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.UpsertStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
					ShiftIds:       []int64{1, 3},
					Decided:        true,
					Lessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				}
				studentOut := &lesson.UpsertStudentShiftsResponse{
					Submission: submission,
					Shifts:     studentShifts,
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().UpsertStudentShifts(gomock.Any(), studentIn).Return(studentOut, nil)
			},
			summaryID: "1",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionResponse{
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
								{ID: 2, Enabled: false, StartTime: "1830", EndTime: "2000"},
							},
						},
						{Date: "20220202", IsClosed: true, Lessons: entity.StudentShifts{}},
						{
							Date:     "20220203",
							IsClosed: false,
							Lessons: entity.StudentShifts{
								{ID: 3, Enabled: true, StartTime: "1700", EndTime: "1830"},
							},
						},
						{Date: "20220204", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220205", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220206", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220207", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220208", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220209", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220210", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220211", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220212", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220213", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220214", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220215", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220216", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220217", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220218", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220219", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220220", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220221", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220222", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220223", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220224", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220225", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220226", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220227", IsClosed: true, Lessons: entity.StudentShifts{}},
						{Date: "20220228", IsClosed: true, Lessons: entity.StudentShifts{}},
					},
					SuggestedLessons: entity.StudentSuggestedLessons{
						{SubjectID: 1, Total: 4},
					},
				},
			},
		},
		{
			name: "failed to invalid summary id",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
			},
			summaryID: "aaa",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
			},
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
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
			},
			summaryID: "1",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
			},
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
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(nil, errmock)
			},
			summaryID: "1",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
			},
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
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
			},
			summaryID: "1",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to upsert student shifts",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				studentIn := &lesson.UpsertStudentShiftsRequest{
					StudentId:      idmock,
					ShiftSummaryId: 1,
					ShiftIds:       []int64{1, 3},
					Decided:        true,
					Lessons: []*lesson.StudentSuggestedLesson{
						{SubjectId: 1, Total: 4},
					},
				}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().UpsertStudentShifts(gomock.Any(), studentIn).Return(nil, errmock)
			},
			summaryID: "1",
			req: &request.UpsertSubmissionRequest{
				ShiftIDs: []int64{1, 3},
				SuggestedLessons: []*request.SuggestedLesson{
					{SubjectID: 1, Total: 4},
				},
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
			path := fmt.Sprintf("/v1/submissions/%s", tt.summaryID)
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestGetSubmissionTemplate(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 12, 0, 0, 0, 0)
	schedules := []*classroom.Schedule{
		{
			Weekday:  int32(time.Sunday),
			IsClosed: true,
			Lessons:  []*classroom.Schedule_Lesson{},
		},
		{
			Weekday:  int32(time.Monday),
			IsClosed: false,
			Lessons: []*classroom.Schedule_Lesson{
				{StartTime: "1700", EndTime: "1830"},
				{StartTime: "1830", EndTime: "2000"},
			},
		},
	}
	template := &lesson.StudentShiftTemplate{
		StudentId: "studentid",
		Schedules: []*lesson.ShiftSchedule{
			{
				Weekday: int32(time.Monday),
				Lessons: []*lesson.LessonSchedule{
					{StartTime: "1700", EndTime: "1830"},
				},
			},
		},
		SuggestedLessons: []*lesson.SuggestedLesson{
			{SubjectId: 1, Total: 4},
		},
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				schedulesIn := &classroom.ListSchedulesRequest{}
				schedulesOut := &classroom.ListSchedulesResponse{Schedules: schedules}
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.classroom.EXPECT().ListSchedules(gomock.Any(), schedulesIn).Return(schedulesOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.SubmissionTemplateResponse{
					Schedules: entity.StudentSchedules{
						{
							Weekday: time.Sunday,
							Lessons: entity.StudentScheduleLessons{},
						},
						{
							Weekday: time.Monday,
							Lessons: entity.StudentScheduleLessons{
								{StartTime: "1700", EndTime: "1830", Enabled: true},
								{StartTime: "1830", EndTime: "2000", Enabled: false},
							},
						},
					},
					SuggestedLessons: entity.StudentSuggestedLessons{
						{SubjectID: 1, Total: 4},
					},
				},
			},
		},
		{
			name: "failed to list schedules",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				schedulesIn := &classroom.ListSchedulesRequest{}
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				templateOut := &lesson.GetStudentShiftTemplateResponse{Template: template}
				mocks.classroom.EXPECT().ListSchedules(gomock.Any(), schedulesIn).Return(nil, errmock)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(templateOut, nil)
			},
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get student shift template",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				schedulesIn := &classroom.ListSchedulesRequest{}
				schedulesOut := &classroom.ListSchedulesResponse{Schedules: schedules}
				templateIn := &lesson.GetStudentShiftTemplateRequest{StudentId: idmock}
				mocks.classroom.EXPECT().ListSchedules(gomock.Any(), schedulesIn).Return(schedulesOut, nil)
				mocks.lesson.EXPECT().GetStudentShiftTemplate(gomock.Any(), templateIn).Return(nil, errmock)
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
			const path = "/v1/me/submission"
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestUpsertSubmissionTemplate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		req    *request.UpsertSubmissionTemplateRequest
		expect *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpsertStudentShiftTemplateRequest{
					StudentId: idmock,
					Template: &lesson.StudentShiftTemplateToUpsert{
						Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
							{
								Weekday: int32(time.Sunday),
								Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{},
							},
							{
								Weekday: int32(time.Monday),
								Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{
									{StartTime: "1700", EndTime: "1830"},
								},
							},
						},
						SuggestedLessons: []*lesson.StudentSuggestedLesson{
							{SubjectId: 1, Total: 4},
						},
					},
				}
				out := &lesson.UpsertStudentShiftTemplateResponse{}
				mocks.lesson.EXPECT().UpsertStudentShiftTemplate(gomock.Any(), in).Return(out, nil)
			},
			req: &request.UpsertSubmissionTemplateRequest{
				Schedules: entity.StudentSchedules{
					{
						Weekday: time.Sunday,
						Lessons: entity.StudentScheduleLessons{},
					},
					{
						Weekday: time.Monday,
						Lessons: entity.StudentScheduleLessons{
							{StartTime: "1700", EndTime: "1830", Enabled: true},
							{StartTime: "1830", EndTime: "2000", Enabled: false},
						},
					},
				},
				SuggestedLessons: entity.StudentSuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
			},
			expect: &testResponse{
				code: http.StatusNoContent,
			},
		},
		{
			name: "failed to upsert student shift template",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				in := &lesson.UpsertStudentShiftTemplateRequest{
					StudentId: idmock,
					Template: &lesson.StudentShiftTemplateToUpsert{
						Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
							{
								Weekday: int32(time.Sunday),
								Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{},
							},
							{
								Weekday: int32(time.Monday),
								Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{
									{StartTime: "1700", EndTime: "1830"},
								},
							},
						},
						SuggestedLessons: []*lesson.StudentSuggestedLesson{
							{SubjectId: 1, Total: 4},
						},
					},
				}
				mocks.lesson.EXPECT().UpsertStudentShiftTemplate(gomock.Any(), in).Return(nil, errmock)
			},
			req: &request.UpsertSubmissionTemplateRequest{
				Schedules: entity.StudentSchedules{
					{
						Weekday: time.Sunday,
						Lessons: entity.StudentScheduleLessons{},
					},
					{
						Weekday: time.Monday,
						Lessons: entity.StudentScheduleLessons{
							{StartTime: "1700", EndTime: "1830", Enabled: true},
							{StartTime: "1830", EndTime: "2000", Enabled: false},
						},
					},
				},
				SuggestedLessons: entity.StudentSuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
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
			const path = "/v1/me/submission"
			req := newHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
