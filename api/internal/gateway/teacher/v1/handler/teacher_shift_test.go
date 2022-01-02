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

func TestListTeacherSubmissions(t *testing.T) {
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
	submissions := []*lesson.TeacherSubmission{
		{
			TeacherId:      idmock,
			ShiftSummaryId: 1,
			Decided:        true,
			CreatedAt:      now.Unix(),
			UpdatedAt:      now.Unix(),
		},
	}

	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		teacherID string
		query     string
		expect    *testResponse
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
				submissionsIn := &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
					TeacherId:       idmock,
					ShiftSummaryIds: []int64{1, 2},
				}
				submissionsOut := &lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse{
					Submissions: submissions,
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().
					ListTeacherSubmissionsByShiftSummaryIDs(gomock.Any(), submissionsIn).
					Return(submissionsOut, nil)
			},
			teacherID: idmock,
			query:     "",
			expect: &testResponse{
				code: http.StatusOK,
				body: &response.TeacherSubmissionsResponse{
					Summaries: entity.TeacherSubmissions{
						{
							ID:               1,
							Year:             2022,
							Month:            2,
							ShiftStatus:      entity.ShiftStatusAccepting,
							SubmissionStatus: entity.TeacherSubmissionStatusSubmitted,
							OpenAt:           jst.Date(2022, 1, 1, 0, 0, 0, 0),
							EndAt:            jst.Date(2022, 1, 15, 0, 0, 0, 0),
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						{
							ID:               2,
							Year:             2022,
							Month:            3,
							ShiftStatus:      entity.ShiftStatusWaiting,
							SubmissionStatus: entity.TeacherSubmissionStatusWaiting,
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
			name:      "failed to parse limit",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			teacherID: idmock,
			query:     "?limit=aaa",
			expect: &testResponse{
				code: http.StatusBadRequest,
			},
		},
		{
			name:      "failed to parse offset",
			setup:     func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {},
			teacherID: idmock,
			query:     "?offset=aaa",
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
			teacherID: idmock,
			query:     "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list teacher submissions",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				shiftsIn := &lesson.ListShiftSummariesRequest{
					Limit:   30,
					Offset:  0,
					Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
				}
				shiftsOut := &lesson.ListShiftSummariesResponse{Summaries: summaries}
				submissionsIn := &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
					TeacherId:       idmock,
					ShiftSummaryIds: []int64{1, 2},
				}
				mocks.lesson.EXPECT().ListShiftSummaries(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
				mocks.lesson.EXPECT().
					ListTeacherSubmissionsByShiftSummaryIDs(gomock.Any(), submissionsIn).
					Return(nil, errmock)
			},
			teacherID: idmock,
			query:     "",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := fmt.Sprintf("/v1/teachers/%s/submissions%s", tt.teacherID, tt.query)
			req := newHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}
