package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/golang/mock/gomock"
)

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
					Shifts: map[string]entity.Shifts{
						"20220201": {
							{
								ID:        1,
								Date:      "20220201",
								StartTime: "1700",
								EndTime:   "1830",
							},
							{
								ID:        2,
								Date:      "20220201",
								StartTime: "1830",
								EndTime:   "2000",
							},
						},
						"20220203": {
							{
								ID:        3,
								Date:      "20220203",
								StartTime: "1700",
								EndTime:   "1830",
							},
						},
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
