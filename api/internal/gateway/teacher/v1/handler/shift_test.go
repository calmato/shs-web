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
					Limit:  30,
					Offset: 0,
					Status: lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
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
					Limit:  30,
					Offset: 0,
					Status: lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
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

	tests := []struct {
		name    string
		setup   func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller)
		shiftID string
		expect  *testResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				summaryOut := &lesson.GetShiftSummaryResponse{Summary: summary}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
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
					Shifts: map[string]*entity.ShiftDetail{
						"20220201": {
							IsClosed: false,
							Lessons: []*entity.ShiftDetailLesson{
								{ID: 1, StartTime: "1700", EndTime: "1830"},
								{ID: 2, StartTime: "1830", EndTime: "2000"},
							},
						},
						"20220202": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220203": {
							IsClosed: false,
							Lessons: []*entity.ShiftDetailLesson{
								{ID: 3, StartTime: "1700", EndTime: "1830"},
							},
						},
						"20220204": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220205": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220206": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220207": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220208": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220209": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220210": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220211": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220212": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220213": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220214": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220215": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220216": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220217": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220218": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220219": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220220": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220221": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220222": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220223": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220224": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220225": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220226": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220227": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220228": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
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
			name: "failed to get shift summary",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
				summaryIn := &lesson.GetShiftSummaryRequest{Id: 1}
				shiftsIn := &lesson.ListShiftsRequest{ShiftSummaryId: 1}
				shiftsOut := &lesson.ListShiftsResponse{Shifts: shifts}
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(nil, errmock)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
			},
			shiftID: "1",
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
			shiftID: "1",
			expect: &testResponse{
				code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to new shift details",
			setup: func(ctx context.Context, t *testing.T, mocks *mocks, ctrl *gomock.Controller) {
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
				mocks.lesson.EXPECT().GetShiftSummary(gomock.Any(), summaryIn).Return(summaryOut, nil)
				mocks.lesson.EXPECT().ListShifts(gomock.Any(), shiftsIn).Return(shiftsOut, nil)
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
					EndAt:       jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
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
					Shifts: map[string]*entity.ShiftDetail{
						"20220201": {
							IsClosed: false,
							Lessons: []*entity.ShiftDetailLesson{
								{ID: 1, StartTime: "1700", EndTime: "1830"},
								{ID: 2, StartTime: "1830", EndTime: "2000"},
							},
						},
						"20220202": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220203": {
							IsClosed: false,
							Lessons: []*entity.ShiftDetailLesson{
								{ID: 3, StartTime: "1700", EndTime: "1830"},
							},
						},
						"20220204": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220205": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220206": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220207": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220208": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220209": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220210": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220211": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220212": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220213": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220214": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220215": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220216": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220217": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220218": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220219": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220220": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220221": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220222": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220223": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220224": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220225": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220226": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220227": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
						"20220228": {IsClosed: true, Lessons: []*entity.ShiftDetailLesson{}},
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
					EndAt:       jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
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
					EndAt:       jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
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
